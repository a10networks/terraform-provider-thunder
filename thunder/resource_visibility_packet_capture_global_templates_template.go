package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template`: Configure global packet capture template for T1 counters\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateDelete,

		Schema: map[string]*schema.Schema{
			"capture_config": {
				Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Template Name",
			},
			"trigger_sys_obj_stats_change": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"system_ctr_lib_acct": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"total_nodes_free_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total nodes free failed",
												},
												"total_nodes_unlink_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total nodes unlink failed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"total_nodes_free_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total nodes free failed",
												},
												"total_nodes_unlink_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total nodes unlink failed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"system_hardware_accelerate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hw_fwd_prog_errors": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hardware forward programming Errors",
												},
												"hw_fwd_flow_singlebit_errors": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hardware forward singlebit Errors",
												},
												"hw_fwd_flow_tag_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hardware forward tag mismatch errors",
												},
												"hw_fwd_flow_seq_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hardware forward sequence mismatch errors",
												},
												"hw_fwd_flow_error_count": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hardware forward flow error count",
												},
												"hw_fwd_flow_unalign_count": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hardware forward flow unalign count",
												},
												"hw_fwd_flow_underflow_count": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hardware forward flow underflow count",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"hw_fwd_prog_errors": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hardware forward programming Errors",
												},
												"hw_fwd_flow_singlebit_errors": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hardware forward singlebit Errors",
												},
												"hw_fwd_flow_tag_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hardware forward tag mismatch errors",
												},
												"hw_fwd_flow_seq_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hardware forward sequence mismatch errors",
												},
												"hw_fwd_flow_error_count": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hardware forward flow error count",
												},
												"hw_fwd_flow_unalign_count": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hardware forward flow unalign count",
												},
												"hw_fwd_flow_underflow_count": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hardware forward flow underflow count",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"system_radius_server": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"radius_request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Dropped (Malformed Packet)",
												},
												"request_bad_secret_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Bad Secret Dropped",
												},
												"request_no_key_vap_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request No Key Attribute Dropped",
												},
												"request_malformed_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Malformed Dropped",
												},
												"radius_table_full": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Dropped (Table Full)",
												},
												"secret_not_configured_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Secret Not Configured Dropped",
												},
												"ha_standby_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA Standby Dropped",
												},
												"ipv6_prefix_length_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Framed IPV6 Prefix Length Mismatch",
												},
												"invalid_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Radius Request has Invalid Key Field",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"radius_request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Dropped (Malformed Packet)",
												},
												"request_bad_secret_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Bad Secret Dropped",
												},
												"request_no_key_vap_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request No Key Attribute Dropped",
												},
												"request_malformed_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Malformed Dropped",
												},
												"radius_table_full": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Dropped (Table Full)",
												},
												"secret_not_configured_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Secret Not Configured Dropped",
												},
												"ha_standby_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA Standby Dropped",
												},
												"ipv6_prefix_length_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Framed IPV6 Prefix Length Mismatch",
												},
												"invalid_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Radius Request has Invalid Key Field",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"system_ip_threat_list": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"error_out_of_memory": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of memory Error",
												},
												"error_out_of_spe_entries": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of SPE Entries Error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"error_out_of_memory": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of memory Error",
												},
												"error_out_of_spe_entries": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of SPE Entries Error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"system_fpga_drop": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mrx_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total MRX Drop",
												},
												"hrx_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total HRX Drop",
												},
												"siz_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Size Drop",
												},
												"fcs_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total FCS Drop",
												},
												"land_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total LAND Attack Drop",
												},
												"empty_frag_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Empty frag Drop",
												},
												"mic_frag_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Micro frag Drop",
												},
												"ipv4_opt_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IPv4 opt Drop",
												},
												"ipv4_frag": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IP frag Drop",
												},
												"bad_ip_hdr_len": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP hdr len Drop",
												},
												"bad_ip_flags_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP Flags Drop",
												},
												"bad_ip_ttl_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP TTL Drop",
												},
												"no_ip_payload_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total No IP Payload Drop",
												},
												"oversize_ip_payload": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Oversize IP PL Drop",
												},
												"bad_ip_payload_len": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP PL len Drop",
												},
												"bad_ip_frag_offset": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP frag off Drop",
												},
												"bad_ip_chksum_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP csum Drop",
												},
												"icmp_pod_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total ICMP POD Drop",
												},
												"tcp_bad_urg_offet": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP bad urg off Drop",
												},
												"tcp_short_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP short hdr Drop",
												},
												"tcp_bad_ip_len": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP Bad IP Len Drop",
												},
												"tcp_null_flags": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP null flags Drop",
												},
												"tcp_null_scan": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP null scan Drop",
												},
												"tcp_fin_sin": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP SYN&FIN Drop",
												},
												"tcp_xmas_flags": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP XMAS FLAGS Drop",
												},
												"tcp_xmas_scan": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP XMAS scan Drop",
												},
												"tcp_syn_frag": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP SYN frag Drop",
												},
												"tcp_frag_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP frag header Drop",
												},
												"tcp_bad_chksum": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP bad csum Drop",
												},
												"udp_short_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP short hdr Drop",
												},
												"udp_bad_ip_len": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP bad leng Drop",
												},
												"udp_kb_frags": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP KB frag Drop",
												},
												"udp_port_lb": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP port LB Drop",
												},
												"udp_bad_chksum": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP bad csum Drop",
												},
												"runt_ip_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Runt IP hdr Drop",
												},
												"runt_tcpudp_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Runt TCPUDP hdr Drop",
												},
												"tun_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Tun mismatch Drop",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"mrx_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total MRX Drop",
												},
												"hrx_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total HRX Drop",
												},
												"siz_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Size Drop",
												},
												"fcs_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total FCS Drop",
												},
												"land_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total LAND Attack Drop",
												},
												"empty_frag_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Empty frag Drop",
												},
												"mic_frag_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Micro frag Drop",
												},
												"ipv4_opt_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IPv4 opt Drop",
												},
												"ipv4_frag": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IP frag Drop",
												},
												"bad_ip_hdr_len": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP hdr len Drop",
												},
												"bad_ip_flags_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP Flags Drop",
												},
												"bad_ip_ttl_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP TTL Drop",
												},
												"no_ip_payload_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total No IP Payload Drop",
												},
												"oversize_ip_payload": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Oversize IP PL Drop",
												},
												"bad_ip_payload_len": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP PL len Drop",
												},
												"bad_ip_frag_offset": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP frag off Drop",
												},
												"bad_ip_chksum_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Bad IP csum Drop",
												},
												"icmp_pod_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total ICMP POD Drop",
												},
												"tcp_bad_urg_offet": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP bad urg off Drop",
												},
												"tcp_short_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP short hdr Drop",
												},
												"tcp_bad_ip_len": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP Bad IP Len Drop",
												},
												"tcp_null_flags": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP null flags Drop",
												},
												"tcp_null_scan": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP null scan Drop",
												},
												"tcp_fin_sin": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP SYN&FIN Drop",
												},
												"tcp_xmas_flags": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP XMAS FLAGS Drop",
												},
												"tcp_xmas_scan": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP XMAS scan Drop",
												},
												"tcp_syn_frag": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP SYN frag Drop",
												},
												"tcp_frag_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP frag header Drop",
												},
												"tcp_bad_chksum": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP bad csum Drop",
												},
												"udp_short_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP short hdr Drop",
												},
												"udp_bad_ip_len": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP bad leng Drop",
												},
												"udp_kb_frags": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP KB frag Drop",
												},
												"udp_port_lb": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP port LB Drop",
												},
												"udp_bad_chksum": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total UDP bad csum Drop",
												},
												"runt_ip_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Runt IP hdr Drop",
												},
												"runt_tcpudp_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Runt TCPUDP hdr Drop",
												},
												"tun_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Tun mismatch Drop",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"system_dpdk_stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"pkt_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total packet drop",
												},
												"pkt_lnk_down_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total packets link down drop",
												},
												"err_pkt_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total error packet drop",
												},
												"rx_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet error",
												},
												"tx_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet error",
												},
												"tx_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet drop",
												},
												"rx_len_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet length error",
												},
												"rx_over_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet over error",
												},
												"rx_crc_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet CRC error",
												},
												"rx_frame_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet frame error",
												},
												"rx_no_buff_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet no buffer error",
												},
												"rx_miss_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet miss error",
												},
												"tx_abort_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet abort error",
												},
												"tx_carrier_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packert carrier error",
												},
												"tx_fifo_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet fifo error",
												},
												"tx_hbeat_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet HBEAT error",
												},
												"tx_windows_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX windows error",
												},
												"rx_long_len_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet long length error",
												},
												"rx_short_len_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet short length error",
												},
												"rx_align_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet align error",
												},
												"rx_csum_offload_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Rx packet check-sum offload error",
												},
												"io_rx_que_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IO core Rx queue drop",
												},
												"io_tx_que_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IO core TX queue drop",
												},
												"io_ring_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IO core ring drop",
												},
												"w_tx_que_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total worker core queue drop",
												},
												"w_link_down_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total worker core link down drop",
												},
												"w_ring_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total worker core ring drop",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"pkt_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total packet drop",
												},
												"pkt_lnk_down_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total packets link down drop",
												},
												"err_pkt_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total error packet drop",
												},
												"rx_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet error",
												},
												"tx_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet error",
												},
												"tx_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet drop",
												},
												"rx_len_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet length error",
												},
												"rx_over_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet over error",
												},
												"rx_crc_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet CRC error",
												},
												"rx_frame_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet frame error",
												},
												"rx_no_buff_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet no buffer error",
												},
												"rx_miss_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet miss error",
												},
												"tx_abort_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet abort error",
												},
												"tx_carrier_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packert carrier error",
												},
												"tx_fifo_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet fifo error",
												},
												"tx_hbeat_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX packet HBEAT error",
												},
												"tx_windows_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TX windows error",
												},
												"rx_long_len_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet long length error",
												},
												"rx_short_len_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet short length error",
												},
												"rx_align_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RX packet align error",
												},
												"rx_csum_offload_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Rx packet check-sum offload error",
												},
												"io_rx_que_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IO core Rx queue drop",
												},
												"io_tx_que_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IO core TX queue drop",
												},
												"io_ring_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total IO core ring drop",
												},
												"w_tx_que_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total worker core queue drop",
												},
												"w_link_down_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total worker core link down drop",
												},
												"w_ring_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total worker core ring drop",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"ip_anomaly_drop": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"land": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Land Attack Drop",
												},
												"emp_frg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Empty Fragment Drop",
												},
												"emp_mic_frg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Micro Fragment Drop",
												},
												"opt": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv4 Options Drop",
												},
												"frg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv4 Fragment Drop",
												},
												"bad_ip_hdrlen": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Header Len Drop",
												},
												"bad_ip_flg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Flags Drop",
												},
												"bad_ip_ttl": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP TTL Drop",
												},
												"no_ip_payload": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No IP Payload drop",
												},
												"over_ip_payload": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Oversize IP Payload Drop",
												},
												"bad_ip_payload_len": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Payload Len Drop",
												},
												"bad_ip_frg_offset": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Fragment Offset Drop",
												},
												"csum": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Checksum Drop",
												},
												"pod": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMP Ping of Death Drop",
												},
												"bad_tcp_urg_offset": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Bad Urgent Offset Drop",
												},
												"tcp_sht_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Short Header Drop",
												},
												"tcp_bad_iplen": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Bad IP Length Drop",
												},
												"tcp_null_frg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Null Flags Drop",
												},
												"tcp_null_scan": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Null Scan Drop",
												},
												"tcp_syn_fin": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Syn and Fin Drop",
												},
												"tcp_xmas": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP XMAS Flags Drop",
												},
												"tcp_xmas_scan": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP XMAS Scan Drop",
												},
												"tcp_syn_frg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Syn Fragment Drop",
												},
												"tcp_frg_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Fragmented Header Drop",
												},
												"tcp_bad_csum": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Bad Checksum Drop",
												},
												"udp_srt_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP Short Header Drop",
												},
												"udp_bad_len": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP Bad Length Drop",
												},
												"udp_kerb_frg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP Kerberos Fragment Drop",
												},
												"udp_port_lb": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP Port Loopback Drop",
												},
												"udp_bad_csum": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP Bad Checksum Drop",
												},
												"runt_ip_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Runt IP Header Drop",
												},
												"runt_tcp_udp_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Runt TCP/UDP Header Drop",
												},
												"ipip_tnl_msmtch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IP-over-IP Tunnel Mismatch Drop",
												},
												"tcp_opt_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Option Error Drop",
												},
												"ipip_tnl_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IP-over-IP Tunnel Error Drop",
												},
												"vxlan_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for VXLAN Tunnel Error Drop",
												},
												"nvgre_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GRE Tunnel Error Drop",
												},
												"gre_pptp_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GRE PPTP Error Drop",
												},
												"ipv6_eh_hbh": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Hop by Hop Header Drop",
												},
												"ipv6_eh_dest": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Destination Header Drop",
												},
												"ipv6_eh_routing": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Routing Header Drop",
												},
												"ipv6_eh_frag": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Fragmentation Header Drop",
												},
												"ipv6_eh_ah": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Authentication Header Drop",
												},
												"ipv6_eh_esp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 ESP Header Drop",
												},
												"ipv6_eh_mobility": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Mobility Header Drop",
												},
												"ipv6_eh_none": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 No Next Header Drop",
												},
												"ipv6_eh_other": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Unknown Extension Header Drop",
												},
												"ipv6_eh_malformed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Malformed Extension Header Drop",
												},
												"tcp_udp_zero_port": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP UDP Zero Port Drop",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"land": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Land Attack Drop",
												},
												"emp_frg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Empty Fragment Drop",
												},
												"emp_mic_frg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Micro Fragment Drop",
												},
												"opt": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv4 Options Drop",
												},
												"frg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv4 Fragment Drop",
												},
												"bad_ip_hdrlen": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Header Len Drop",
												},
												"bad_ip_flg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Flags Drop",
												},
												"bad_ip_ttl": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP TTL Drop",
												},
												"no_ip_payload": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No IP Payload drop",
												},
												"over_ip_payload": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Oversize IP Payload Drop",
												},
												"bad_ip_payload_len": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Payload Len Drop",
												},
												"bad_ip_frg_offset": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Fragment Offset Drop",
												},
												"csum": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad IP Checksum Drop",
												},
												"pod": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMP Ping of Death Drop",
												},
												"bad_tcp_urg_offset": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Bad Urgent Offset Drop",
												},
												"tcp_sht_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Short Header Drop",
												},
												"tcp_bad_iplen": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Bad IP Length Drop",
												},
												"tcp_null_frg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Null Flags Drop",
												},
												"tcp_null_scan": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Null Scan Drop",
												},
												"tcp_syn_fin": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Syn and Fin Drop",
												},
												"tcp_xmas": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP XMAS Flags Drop",
												},
												"tcp_xmas_scan": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP XMAS Scan Drop",
												},
												"tcp_syn_frg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Syn Fragment Drop",
												},
												"tcp_frg_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Fragmented Header Drop",
												},
												"tcp_bad_csum": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Bad Checksum Drop",
												},
												"udp_srt_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP Short Header Drop",
												},
												"udp_bad_len": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP Bad Length Drop",
												},
												"udp_kerb_frg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP Kerberos Fragment Drop",
												},
												"udp_port_lb": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP Port Loopback Drop",
												},
												"udp_bad_csum": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP Bad Checksum Drop",
												},
												"runt_ip_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Runt IP Header Drop",
												},
												"runt_tcp_udp_hdr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Runt TCP/UDP Header Drop",
												},
												"ipip_tnl_msmtch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IP-over-IP Tunnel Mismatch Drop",
												},
												"tcp_opt_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Option Error Drop",
												},
												"ipip_tnl_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IP-over-IP Tunnel Error Drop",
												},
												"vxlan_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for VXLAN Tunnel Error Drop",
												},
												"nvgre_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GRE Tunnel Error Drop",
												},
												"gre_pptp_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GRE PPTP Error Drop",
												},
												"ipv6_eh_hbh": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Hop by Hop Header Drop",
												},
												"ipv6_eh_dest": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Destination Header Drop",
												},
												"ipv6_eh_routing": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Routing Header Drop",
												},
												"ipv6_eh_frag": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Fragmentation Header Drop",
												},
												"ipv6_eh_ah": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Authentication Header Drop",
												},
												"ipv6_eh_esp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 ESP Header Drop",
												},
												"ipv6_eh_mobility": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Mobility Header Drop",
												},
												"ipv6_eh_none": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 No Next Header Drop",
												},
												"ipv6_eh_other": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Unknown Extension Header Drop",
												},
												"ipv6_eh_malformed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for IPv6 Malformed Extension Header Drop",
												},
												"tcp_udp_zero_port": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP UDP Zero Port Drop",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"aam_authentication_global": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"misses": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Authentication Request Missed",
												},
												"open_socket_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total AAM Open Socket Failed",
												},
												"connect_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total AAM Connect Failed",
												},
												"create_timer_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total AAM Timer Creation Failed",
												},
												"get_socket_option_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total AAM Get Socket Option Failed",
												},
												"aflex_authz_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Authorization failure number in aFleX",
												},
												"authn_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Authentication failure number",
												},
												"authz_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Authorization failure number",
												},
												"dns_resolve_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total AAM DNS resolve failed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"misses": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Authentication Request Missed",
												},
												"open_socket_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total AAM Open Socket Failed",
												},
												"connect_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total AAM Connect Failed",
												},
												"create_timer_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total AAM Timer Creation Failed",
												},
												"get_socket_option_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total AAM Get Socket Option Failed",
												},
												"aflex_authz_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Authorization failure number in aFleX",
												},
												"authn_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Authentication failure number",
												},
												"authz_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Authorization failure number",
												},
												"dns_resolve_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total AAM DNS resolve failed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"aam_rdns": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Dropped Request",
												},
												"response_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Failure Response",
												},
												"response_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Error Response",
												},
												"response_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout Response",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Dropped Request",
												},
												"response_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Failure Response",
												},
												"response_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Error Response",
												},
												"response_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout Response",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"aam_auth_server_ldap": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"admin_bind_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Admin Bind Failure",
												},
												"bind_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total User Bind Failure",
												},
												"search_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Search Failure",
												},
												"authorize_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Authorization Failure",
												},
												"timeout_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout",
												},
												"other_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Other Error",
												},
												"request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Dropped Request",
												},
												"response_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Failure Response",
												},
												"response_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Error Response",
												},
												"response_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout Response",
												},
												"job_start_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Job Start Error",
												},
												"polling_control_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Polling Control Error",
												},
												"ssl_session_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TLS/SSL Session Failure",
												},
												"pw_change_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total password change failure",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"admin_bind_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Admin Bind Failure",
												},
												"bind_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total User Bind Failure",
												},
												"search_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Search Failure",
												},
												"authorize_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Authorization Failure",
												},
												"timeout_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout",
												},
												"other_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Other Error",
												},
												"request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Dropped Request",
												},
												"response_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Failure Response",
												},
												"response_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Error Response",
												},
												"response_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout Response",
												},
												"job_start_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Job Start Error",
												},
												"polling_control_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Polling Control Error",
												},
												"ssl_session_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TLS/SSL Session Failure",
												},
												"pw_change_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total password change failure",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"aam_auth_server_ocsp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"stapling_request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Stapling Dropped Request",
												},
												"stapling_response_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Stapling Failure Response",
												},
												"stapling_response_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Stapling Error Response",
												},
												"stapling_response_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Stapling Timeout Response",
												},
												"request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Dropped Request",
												},
												"response_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Failure Response",
												},
												"response_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Error Response",
												},
												"response_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Timeout Response",
												},
												"job_start_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Job Start Error",
												},
												"polling_control_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Polling Control Error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"stapling_request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Stapling Dropped Request",
												},
												"stapling_response_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Stapling Failure Response",
												},
												"stapling_response_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Stapling Error Response",
												},
												"stapling_response_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Stapling Timeout Response",
												},
												"request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Dropped Request",
												},
												"response_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Failure Response",
												},
												"response_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Error Response",
												},
												"response_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Timeout Response",
												},
												"job_start_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Job Start Error",
												},
												"polling_control_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total OCSP Polling Control Error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"aam_auth_server_radius": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"authen_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Authentication Failure",
												},
												"authorize_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Authorization Failure",
												},
												"timeout_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout",
												},
												"other_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Other Error",
												},
												"request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Dropped Request",
												},
												"response_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Failure Response",
												},
												"response_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Error Response",
												},
												"response_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout Response",
												},
												"job_start_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Job Start Error",
												},
												"polling_control_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Polling Control Error",
												},
												"accounting_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Accounting Failure",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"authen_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Authentication Failure",
												},
												"authorize_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Authorization Failure",
												},
												"timeout_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout",
												},
												"other_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Other Error",
												},
												"request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Dropped Request",
												},
												"response_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Failure Response",
												},
												"response_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Error Response",
												},
												"response_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout Response",
												},
												"job_start_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Job Start Error",
												},
												"polling_control_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Polling Control Error",
												},
												"accounting_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Accounting Failure",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"aam_auth_server_win": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"kerberos_timeout_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Timeout",
												},
												"kerberos_other_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Other Error",
												},
												"kerberos_request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Dropped Request",
												},
												"kerberos_response_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Failure Response",
												},
												"kerberos_response_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Error Response",
												},
												"kerberos_response_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Timeout Response",
												},
												"kerberos_job_start_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Job Start Error",
												},
												"kerberos_polling_control_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Polling Control Error",
												},
												"kerberos_pw_expiry": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos password expiry",
												},
												"kerberos_pw_change_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos password change failure",
												},
												"kerberos_validate_kdc_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos KDC Validation Failure",
												},
												"kerberos_generate_kdc_keytab_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos KDC Keytab Generation Failure",
												},
												"kerberos_delete_kdc_keytab_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos KDC Keytab Deletion Failure",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"kerberos_timeout_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Timeout",
												},
												"kerberos_other_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Other Error",
												},
												"kerberos_request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Dropped Request",
												},
												"kerberos_response_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Failure Response",
												},
												"kerberos_response_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Error Response",
												},
												"kerberos_response_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Timeout Response",
												},
												"kerberos_job_start_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Job Start Error",
												},
												"kerberos_polling_control_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Polling Control Error",
												},
												"kerberos_pw_expiry": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos password expiry",
												},
												"kerberos_pw_change_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos password change failure",
												},
												"kerberos_validate_kdc_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos KDC Validation Failure",
												},
												"kerberos_generate_kdc_keytab_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos KDC Keytab Generation Failure",
												},
												"kerberos_delete_kdc_keytab_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos KDC Keytab Deletion Failure",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"aam_auth_account": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Dropped Request",
												},
												"response_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Failure Response",
												},
												"response_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Error Response",
												},
												"response_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout Response",
												},
												"response_other": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Other Response",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Dropped Request",
												},
												"response_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Failure Response",
												},
												"response_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Error Response",
												},
												"response_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout Response",
												},
												"response_other": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Other Response",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"aam_auth_saml_global": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"acs_authz_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total SAML Single-Sign-On Authorization Fail",
												},
												"acs_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total SAML Single-Sign-On Error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"acs_authz_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total SAML Single-Sign-On Authorization Fail",
												},
												"acs_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total SAML Single-Sign-On Error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"aam_auth_relay_kerberos": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"timeout_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout",
												},
												"other_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Other Error",
												},
												"request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Dropped Request",
												},
												"response_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Failure Response",
												},
												"response_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Error Response",
												},
												"response_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout Response",
												},
												"job_start_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Job Start Error",
												},
												"polling_control_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Polling Control Error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"timeout_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout",
												},
												"other_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Other Error",
												},
												"request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Dropped Request",
												},
												"response_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Failure Response",
												},
												"response_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Error Response",
												},
												"response_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout Response",
												},
												"job_start_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Job Start Error",
												},
												"polling_control_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Polling Control Error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"aam_auth_captcha": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Dropped Request",
												},
												"response_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Failure Response",
												},
												"response_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Error Response",
												},
												"response_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout Response",
												},
												"json_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Failure JSON Response",
												},
												"attr_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Attribute Check Failure",
												},
												"timeout_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout",
												},
												"other_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Other Error",
												},
												"job_start_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Job Start Error",
												},
												"polling_control_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Polling Control Error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Dropped Request",
												},
												"response_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Failure Response",
												},
												"response_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Error Response",
												},
												"response_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout Response",
												},
												"json_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Failure JSON Response",
												},
												"attr_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Attribute Check Failure",
												},
												"timeout_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout",
												},
												"other_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Other Error",
												},
												"job_start_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Job Start Error",
												},
												"polling_control_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Polling Control Error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_ssl_error": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"app_data_in_handshake": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for app data in handshake",
												},
												"attempt_to_reuse_sess_in_diff_context": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for attempt to reuse sess in diff context",
												},
												"bad_alert_record": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad alert record",
												},
												"bad_authentication_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad authentication type",
												},
												"bad_change_cipher_spec": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad change cipher spec",
												},
												"bad_checksum": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad checksum",
												},
												"bad_data_returned_by_callback": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad data returned by callback",
												},
												"bad_decompression": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad decompression",
												},
												"bad_dh_g_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad dh g length",
												},
												"bad_dh_pub_key_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad dh pub key length",
												},
												"bad_dh_p_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad dh p length",
												},
												"bad_digest_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad digest length",
												},
												"bad_dsa_signature": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad dsa signature",
												},
												"bad_hello_request": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad hello request",
												},
												"bad_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad length",
												},
												"bad_mac_decode": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad mac decode",
												},
												"bad_message_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad message type",
												},
												"bad_packet_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad packet length",
												},
												"bad_protocol_version_counter": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad protocol version counter",
												},
												"bad_response_argument": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad response argument",
												},
												"bad_rsa_decrypt": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad rsa decrypt",
												},
												"bad_rsa_encrypt": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad rsa encrypt",
												},
												"bad_rsa_e_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad rsa e length",
												},
												"bad_rsa_modulus_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad rsa modulus length",
												},
												"bad_rsa_signature": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad rsa signature",
												},
												"bad_signature": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad signature",
												},
												"bad_ssl_filetype": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad ssl filetype",
												},
												"bad_ssl_session_id_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad ssl session id length",
												},
												"bad_state": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad state",
												},
												"bad_write_retry": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad write retry",
												},
												"bio_not_set": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bio not set",
												},
												"block_cipher_pad_is_wrong": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for block cipher pad is wrong",
												},
												"bn_lib": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bn lib",
												},
												"ca_dn_length_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ca dn length mismatch",
												},
												"ca_dn_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ca dn too long",
												},
												"ccs_received_early": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ccs received early",
												},
												"certificate_verify_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for certificate verify failed",
												},
												"cert_length_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cert length mismatch",
												},
												"challenge_is_different": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for challenge is different",
												},
												"cipher_code_wrong_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cipher code wrong length",
												},
												"cipher_or_hash_unavailable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cipher or hash unavailable",
												},
												"cipher_table_src_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cipher table src error",
												},
												"compressed_length_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for compressed length too long",
												},
												"compression_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for compression failure",
												},
												"compression_library_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for compression library error",
												},
												"connection_id_is_different": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for connection id is different",
												},
												"connection_type_not_set": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for connection type not set",
												},
												"data_between_ccs_and_finished": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data between ccs and finished",
												},
												"data_length_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data length too long",
												},
												"decryption_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for decryption failed",
												},
												"decryption_failed_or_bad_record_mac": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for decryption failed or bad record mac",
												},
												"dh_public_value_length_is_wrong": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for dh public value length is wrong",
												},
												"digest_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for digest check failed",
												},
												"encrypted_length_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for encrypted length too long",
												},
												"error_generating_tmp_rsa_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for error generating tmp rsa key",
												},
												"error_in_received_cipher_list": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for error in received cipher list",
												},
												"excessive_message_size": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for excessive message size",
												},
												"extra_data_in_message": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for extra data in message",
												},
												"got_a_fin_before_a_ccs": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for got a fin before a ccs",
												},
												"https_proxy_request": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for https proxy request",
												},
												"http_request": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for http request",
												},
												"illegal_padding": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for illegal padding",
												},
												"inappropriate_fallback": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for inappropriate fallback",
												},
												"invalid_challenge_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid challenge length",
												},
												"invalid_command": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid command",
												},
												"invalid_purpose": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid purpose",
												},
												"invalid_status_response": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid status response",
												},
												"invalid_trust": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid trust",
												},
												"key_arg_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for key arg too long",
												},
												"krb5": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5",
												},
												"krb5_client_cc_principal": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 client cc principal",
												},
												"krb5_client_get_cred": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 client get cred",
												},
												"krb5_client_init": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 client init",
												},
												"krb5_client_mk_req": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 client mk_req",
												},
												"krb5_server_bad_ticket": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 server bad ticket",
												},
												"krb5_server_init": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 server init",
												},
												"krb5_server_rd_req": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 server rd_req",
												},
												"krb5_server_tkt_expired": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 server tkt expired",
												},
												"krb5_server_tkt_not_yet_valid": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 server tkt not yet valid",
												},
												"krb5_server_tkt_skew": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 server tkt skew",
												},
												"length_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for length mismatch",
												},
												"length_too_short": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for length too short",
												},
												"library_bug": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for library bug",
												},
												"library_has_no_ciphers": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for library has no ciphers",
												},
												"mast_key_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for mast key too long",
												},
												"message_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for message too long",
												},
												"missing_dh_dsa_cert": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing dh dsa cert",
												},
												"missing_dh_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing dh key",
												},
												"missing_dh_rsa_cert": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing dh rsa cert",
												},
												"missing_dsa_signing_cert": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing dsa signing cert",
												},
												"missing_export_tmp_dh_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing export tmp dh key",
												},
												"missing_export_tmp_rsa_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing export tmp rsa key",
												},
												"missing_rsa_certificate": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing rsa certificate",
												},
												"missing_rsa_encrypting_cert": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing rsa encrypting cert",
												},
												"missing_rsa_signing_cert": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing rsa signing cert",
												},
												"missing_tmp_dh_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing tmp dh key",
												},
												"missing_tmp_rsa_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing tmp rsa key",
												},
												"missing_tmp_rsa_pkey": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing tmp rsa pkey",
												},
												"missing_verify_message": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing verify message",
												},
												"non_sslv2_initial_packet": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for non sslv2 initial packet",
												},
												"no_certificates_returned": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no certificates returned",
												},
												"no_certificate_assigned": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no certificate assigned",
												},
												"no_certificate_returned": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no certificate returned",
												},
												"no_certificate_set": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no certificate set",
												},
												"no_certificate_specified": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no certificate specified",
												},
												"no_ciphers_available": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no ciphers available",
												},
												"no_ciphers_passed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no ciphers passed",
												},
												"no_ciphers_specified": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no ciphers specified",
												},
												"no_cipher_list": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no cipher list",
												},
												"no_cipher_match": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no cipher match",
												},
												"no_client_cert_received": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no client cert received",
												},
												"no_compression_specified": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no compression specified",
												},
												"no_method_specified": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no method specified",
												},
												"no_privatekey": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no privatekey",
												},
												"no_private_key_assigned": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no private key assigned",
												},
												"no_protocols_available": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no protocols available",
												},
												"no_publickey": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no publickey",
												},
												"no_shared_cipher": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no shared cipher",
												},
												"no_verify_callback": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no verify callback",
												},
												"null_ssl_ctx": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for null ssl ctx",
												},
												"null_ssl_method_passed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for null ssl method passed",
												},
												"old_session_cipher_not_returned": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for old session cipher not returned",
												},
												"packet_length_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for packet length too long",
												},
												"path_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for path too long",
												},
												"peer_did_not_return_a_certificate": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for peer did not return a certificate",
												},
												"peer_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for peer error",
												},
												"peer_error_certificate": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for peer error certificate",
												},
												"peer_error_no_certificate": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for peer error no certificate",
												},
												"peer_error_no_cipher": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for peer error no cipher",
												},
												"peer_error_unsupported_certificate_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for peer error unsupported certificate type",
												},
												"pre_mac_length_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for pre mac length too long",
												},
												"problems_mapping_cipher_functions": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for problems mapping cipher functions",
												},
												"protocol_is_shutdown": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for protocol is shutdown",
												},
												"public_key_encrypt_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for public key encrypt error",
												},
												"public_key_is_not_rsa": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for public key is not rsa",
												},
												"public_key_not_rsa": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for public key not rsa",
												},
												"read_bio_not_set": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for read bio not set",
												},
												"read_wrong_packet_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for read wrong packet type",
												},
												"record_length_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for record length mismatch",
												},
												"record_too_large": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for record too large",
												},
												"record_too_small": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for record too small",
												},
												"required_cipher_missing": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for required cipher missing",
												},
												"reuse_cert_length_not_zero": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for reuse cert length not zero",
												},
												"reuse_cert_type_not_zero": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for reuse cert type not zero",
												},
												"reuse_cipher_list_not_zero": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for reuse cipher list not zero",
												},
												"scsv_received_when_renegotiating": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for scsv received when renegotiating",
												},
												"session_id_context_uninitialized": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for session id context uninitialized",
												},
												"short_read": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for short read",
												},
												"signature_for_non_signing_certificate": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for signature for non signing certificate",
												},
												"ssl23_doing_session_id_reuse": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl23 doing session id reuse",
												},
												"ssl2_connection_id_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl2 connection id too long",
												},
												"ssl3_session_id_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl3 session id too long",
												},
												"ssl3_session_id_too_short": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl3 session id too short",
												},
												"sslv3_alert_bad_certificate": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert bad certificate",
												},
												"sslv3_alert_bad_record_mac": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert bad record mac",
												},
												"sslv3_alert_certificate_expired": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert certificate expired",
												},
												"sslv3_alert_certificate_revoked": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert certificate revoked",
												},
												"sslv3_alert_certificate_unknown": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert certificate unknown",
												},
												"sslv3_alert_decompression_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert decompression failure",
												},
												"sslv3_alert_handshake_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert handshake failure",
												},
												"sslv3_alert_illegal_parameter": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert illegal parameter",
												},
												"sslv3_alert_no_certificate": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert no certificate",
												},
												"sslv3_alert_peer_error_cert": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert peer error cert",
												},
												"sslv3_alert_peer_error_no_cert": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert peer error no cert",
												},
												"sslv3_alert_peer_error_no_cipher": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert peer error no cipher",
												},
												"sslv3_alert_peer_error_unsupp_cert_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert peer error unsupp cert type",
												},
												"sslv3_alert_unexpected_msg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert unexpected msg",
												},
												"sslv3_alert_unknown_remote_err_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert unknown remote err type",
												},
												"sslv3_alert_unspported_cert": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert unspported cert",
												},
												"ssl_ctx_has_no_default_ssl_version": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl ctx has no default ssl version",
												},
												"ssl_handshake_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl handshake failure",
												},
												"ssl_library_has_no_ciphers": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl library has no ciphers",
												},
												"ssl_session_id_callback_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl session id callback failed",
												},
												"ssl_session_id_conflict": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl session id conflict",
												},
												"ssl_session_id_context_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl session id context too long",
												},
												"ssl_session_id_has_bad_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl session id has bad length",
												},
												"ssl_session_id_is_different": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl session id is different",
												},
												"tlsv1_alert_access_denied": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert access denied",
												},
												"tlsv1_alert_decode_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert decode error",
												},
												"tlsv1_alert_decryption_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert decryption failed",
												},
												"tlsv1_alert_decrypt_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert decrypt error",
												},
												"tlsv1_alert_export_restriction": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert export restriction",
												},
												"tlsv1_alert_insufficient_security": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert insufficient security",
												},
												"tlsv1_alert_internal_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert internal error",
												},
												"tlsv1_alert_no_renegotiation": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert no renegotiation",
												},
												"tlsv1_alert_protocol_version": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert protocol version",
												},
												"tlsv1_alert_record_overflow": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert record overflow",
												},
												"tlsv1_alert_unknown_ca": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert unknown ca",
												},
												"tlsv1_alert_user_cancelled": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert user cancelled",
												},
												"tls_client_cert_req_with_anon_cipher": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tls client cert req with anon cipher",
												},
												"tls_peer_did_not_respond_with_cert_list": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tls peer did not respond with cert list",
												},
												"tls_rsa_encrypted_value_length_is_wrong": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tls rsa encrypted value length is wrong",
												},
												"tried_to_use_unsupported_cipher": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tried to use unsupported cipher",
												},
												"unable_to_decode_dh_certs": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to decode dh certs",
												},
												"unable_to_extract_public_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to extract public key",
												},
												"unable_to_find_dh_parameters": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to find dh parameters",
												},
												"unable_to_find_public_key_parameters": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to find public key parameters",
												},
												"unable_to_find_ssl_method": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to find ssl method",
												},
												"unable_to_load_ssl2_md5_routines": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to load ssl2 md5 routines",
												},
												"unable_to_load_ssl3_md5_routines": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to load ssl3 md5 routines",
												},
												"unable_to_load_ssl3_sha1_routines": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to load ssl3 sha1 routines",
												},
												"unexpected_message": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unexpected message",
												},
												"unexpected_record": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unexpected record",
												},
												"uninitialized": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for uninitialized",
												},
												"unknown_alert_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown alert type",
												},
												"unknown_certificate_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown certificate type",
												},
												"unknown_cipher_returned": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown cipher returned",
												},
												"unknown_cipher_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown cipher type",
												},
												"unknown_key_exchange_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown key exchange type",
												},
												"unknown_pkey_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown pkey type",
												},
												"unknown_protocol": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown protocol",
												},
												"unknown_remote_error_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown remote error type",
												},
												"unknown_ssl_version": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown ssl version",
												},
												"unknown_state": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown state",
												},
												"unsupported_cipher": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported cipher",
												},
												"unsupported_compression_algorithm": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported compression algorithm",
												},
												"unsupported_option": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported option",
												},
												"unsupported_protocol": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported protocol",
												},
												"unsupported_ssl_version": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported ssl version",
												},
												"unsupported_status_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported status type",
												},
												"write_bio_not_set": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for write bio not set",
												},
												"wrong_cipher_returned": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong cipher returned",
												},
												"wrong_message_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong message type",
												},
												"wrong_counter_of_key_bits": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong counter of key bits",
												},
												"wrong_signature_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong signature length",
												},
												"wrong_signature_size": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong signature size",
												},
												"wrong_ssl_version": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong ssl version",
												},
												"wrong_version_counter": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong version counter",
												},
												"x509_lib": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for x509 lib",
												},
												"x509_verification_setup_problems": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for x509 verification setup problems",
												},
												"clienthello_tlsext": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for clienthello tlsext",
												},
												"parse_tlsext": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for parse tlsext",
												},
												"serverhello_tlsext": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for serverhello tlsext",
												},
												"ssl3_ext_invalid_servername": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl3 ext invalid servername",
												},
												"ssl3_ext_invalid_servername_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl3 ext invalid servername type",
												},
												"multiple_sgc_restarts": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for multiple sgc restarts",
												},
												"tls_invalid_ecpointformat_list": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tls invalid ecpointformat list",
												},
												"bad_ecc_cert": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad ecc cert",
												},
												"bad_ecdsa_sig": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad ecdsa sig",
												},
												"bad_ecpoint": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad ecpoint",
												},
												"cookie_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cookie mismatch",
												},
												"unsupported_elliptic_curve": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported elliptic curve",
												},
												"no_required_digest": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no required digest",
												},
												"unsupported_digest_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported digest type",
												},
												"bad_handshake_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad handshake length",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"app_data_in_handshake": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for app data in handshake",
												},
												"attempt_to_reuse_sess_in_diff_context": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for attempt to reuse sess in diff context",
												},
												"bad_alert_record": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad alert record",
												},
												"bad_authentication_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad authentication type",
												},
												"bad_change_cipher_spec": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad change cipher spec",
												},
												"bad_checksum": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad checksum",
												},
												"bad_data_returned_by_callback": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad data returned by callback",
												},
												"bad_decompression": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad decompression",
												},
												"bad_dh_g_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad dh g length",
												},
												"bad_dh_pub_key_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad dh pub key length",
												},
												"bad_dh_p_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad dh p length",
												},
												"bad_digest_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad digest length",
												},
												"bad_dsa_signature": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad dsa signature",
												},
												"bad_hello_request": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad hello request",
												},
												"bad_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad length",
												},
												"bad_mac_decode": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad mac decode",
												},
												"bad_message_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad message type",
												},
												"bad_packet_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad packet length",
												},
												"bad_protocol_version_counter": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad protocol version counter",
												},
												"bad_response_argument": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad response argument",
												},
												"bad_rsa_decrypt": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad rsa decrypt",
												},
												"bad_rsa_encrypt": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad rsa encrypt",
												},
												"bad_rsa_e_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad rsa e length",
												},
												"bad_rsa_modulus_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad rsa modulus length",
												},
												"bad_rsa_signature": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad rsa signature",
												},
												"bad_signature": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad signature",
												},
												"bad_ssl_filetype": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad ssl filetype",
												},
												"bad_ssl_session_id_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad ssl session id length",
												},
												"bad_state": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad state",
												},
												"bad_write_retry": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad write retry",
												},
												"bio_not_set": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bio not set",
												},
												"block_cipher_pad_is_wrong": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for block cipher pad is wrong",
												},
												"bn_lib": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bn lib",
												},
												"ca_dn_length_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ca dn length mismatch",
												},
												"ca_dn_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ca dn too long",
												},
												"ccs_received_early": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ccs received early",
												},
												"certificate_verify_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for certificate verify failed",
												},
												"cert_length_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cert length mismatch",
												},
												"challenge_is_different": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for challenge is different",
												},
												"cipher_code_wrong_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cipher code wrong length",
												},
												"cipher_or_hash_unavailable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cipher or hash unavailable",
												},
												"cipher_table_src_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cipher table src error",
												},
												"compressed_length_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for compressed length too long",
												},
												"compression_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for compression failure",
												},
												"compression_library_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for compression library error",
												},
												"connection_id_is_different": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for connection id is different",
												},
												"connection_type_not_set": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for connection type not set",
												},
												"data_between_ccs_and_finished": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data between ccs and finished",
												},
												"data_length_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data length too long",
												},
												"decryption_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for decryption failed",
												},
												"decryption_failed_or_bad_record_mac": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for decryption failed or bad record mac",
												},
												"dh_public_value_length_is_wrong": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for dh public value length is wrong",
												},
												"digest_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for digest check failed",
												},
												"encrypted_length_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for encrypted length too long",
												},
												"error_generating_tmp_rsa_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for error generating tmp rsa key",
												},
												"error_in_received_cipher_list": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for error in received cipher list",
												},
												"excessive_message_size": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for excessive message size",
												},
												"extra_data_in_message": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for extra data in message",
												},
												"got_a_fin_before_a_ccs": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for got a fin before a ccs",
												},
												"https_proxy_request": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for https proxy request",
												},
												"http_request": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for http request",
												},
												"illegal_padding": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for illegal padding",
												},
												"inappropriate_fallback": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for inappropriate fallback",
												},
												"invalid_challenge_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid challenge length",
												},
												"invalid_command": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid command",
												},
												"invalid_purpose": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid purpose",
												},
												"invalid_status_response": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid status response",
												},
												"invalid_trust": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid trust",
												},
												"key_arg_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for key arg too long",
												},
												"krb5": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5",
												},
												"krb5_client_cc_principal": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 client cc principal",
												},
												"krb5_client_get_cred": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 client get cred",
												},
												"krb5_client_init": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 client init",
												},
												"krb5_client_mk_req": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 client mk_req",
												},
												"krb5_server_bad_ticket": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 server bad ticket",
												},
												"krb5_server_init": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 server init",
												},
												"krb5_server_rd_req": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 server rd_req",
												},
												"krb5_server_tkt_expired": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 server tkt expired",
												},
												"krb5_server_tkt_not_yet_valid": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 server tkt not yet valid",
												},
												"krb5_server_tkt_skew": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for krb5 server tkt skew",
												},
												"length_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for length mismatch",
												},
												"length_too_short": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for length too short",
												},
												"library_bug": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for library bug",
												},
												"library_has_no_ciphers": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for library has no ciphers",
												},
												"mast_key_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for mast key too long",
												},
												"message_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for message too long",
												},
												"missing_dh_dsa_cert": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing dh dsa cert",
												},
												"missing_dh_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing dh key",
												},
												"missing_dh_rsa_cert": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing dh rsa cert",
												},
												"missing_dsa_signing_cert": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing dsa signing cert",
												},
												"missing_export_tmp_dh_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing export tmp dh key",
												},
												"missing_export_tmp_rsa_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing export tmp rsa key",
												},
												"missing_rsa_certificate": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing rsa certificate",
												},
												"missing_rsa_encrypting_cert": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing rsa encrypting cert",
												},
												"missing_rsa_signing_cert": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing rsa signing cert",
												},
												"missing_tmp_dh_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing tmp dh key",
												},
												"missing_tmp_rsa_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing tmp rsa key",
												},
												"missing_tmp_rsa_pkey": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing tmp rsa pkey",
												},
												"missing_verify_message": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for missing verify message",
												},
												"non_sslv2_initial_packet": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for non sslv2 initial packet",
												},
												"no_certificates_returned": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no certificates returned",
												},
												"no_certificate_assigned": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no certificate assigned",
												},
												"no_certificate_returned": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no certificate returned",
												},
												"no_certificate_set": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no certificate set",
												},
												"no_certificate_specified": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no certificate specified",
												},
												"no_ciphers_available": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no ciphers available",
												},
												"no_ciphers_passed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no ciphers passed",
												},
												"no_ciphers_specified": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no ciphers specified",
												},
												"no_cipher_list": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no cipher list",
												},
												"no_cipher_match": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no cipher match",
												},
												"no_client_cert_received": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no client cert received",
												},
												"no_compression_specified": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no compression specified",
												},
												"no_method_specified": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no method specified",
												},
												"no_privatekey": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no privatekey",
												},
												"no_private_key_assigned": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no private key assigned",
												},
												"no_protocols_available": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no protocols available",
												},
												"no_publickey": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no publickey",
												},
												"no_shared_cipher": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no shared cipher",
												},
												"no_verify_callback": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no verify callback",
												},
												"null_ssl_ctx": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for null ssl ctx",
												},
												"null_ssl_method_passed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for null ssl method passed",
												},
												"old_session_cipher_not_returned": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for old session cipher not returned",
												},
												"packet_length_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for packet length too long",
												},
												"path_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for path too long",
												},
												"peer_did_not_return_a_certificate": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for peer did not return a certificate",
												},
												"peer_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for peer error",
												},
												"peer_error_certificate": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for peer error certificate",
												},
												"peer_error_no_certificate": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for peer error no certificate",
												},
												"peer_error_no_cipher": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for peer error no cipher",
												},
												"peer_error_unsupported_certificate_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for peer error unsupported certificate type",
												},
												"pre_mac_length_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for pre mac length too long",
												},
												"problems_mapping_cipher_functions": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for problems mapping cipher functions",
												},
												"protocol_is_shutdown": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for protocol is shutdown",
												},
												"public_key_encrypt_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for public key encrypt error",
												},
												"public_key_is_not_rsa": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for public key is not rsa",
												},
												"public_key_not_rsa": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for public key not rsa",
												},
												"read_bio_not_set": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for read bio not set",
												},
												"read_wrong_packet_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for read wrong packet type",
												},
												"record_length_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for record length mismatch",
												},
												"record_too_large": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for record too large",
												},
												"record_too_small": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for record too small",
												},
												"required_cipher_missing": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for required cipher missing",
												},
												"reuse_cert_length_not_zero": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for reuse cert length not zero",
												},
												"reuse_cert_type_not_zero": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for reuse cert type not zero",
												},
												"reuse_cipher_list_not_zero": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for reuse cipher list not zero",
												},
												"scsv_received_when_renegotiating": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for scsv received when renegotiating",
												},
												"session_id_context_uninitialized": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for session id context uninitialized",
												},
												"short_read": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for short read",
												},
												"signature_for_non_signing_certificate": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for signature for non signing certificate",
												},
												"ssl23_doing_session_id_reuse": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl23 doing session id reuse",
												},
												"ssl2_connection_id_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl2 connection id too long",
												},
												"ssl3_session_id_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl3 session id too long",
												},
												"ssl3_session_id_too_short": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl3 session id too short",
												},
												"sslv3_alert_bad_certificate": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert bad certificate",
												},
												"sslv3_alert_bad_record_mac": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert bad record mac",
												},
												"sslv3_alert_certificate_expired": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert certificate expired",
												},
												"sslv3_alert_certificate_revoked": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert certificate revoked",
												},
												"sslv3_alert_certificate_unknown": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert certificate unknown",
												},
												"sslv3_alert_decompression_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert decompression failure",
												},
												"sslv3_alert_handshake_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert handshake failure",
												},
												"sslv3_alert_illegal_parameter": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert illegal parameter",
												},
												"sslv3_alert_no_certificate": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert no certificate",
												},
												"sslv3_alert_peer_error_cert": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert peer error cert",
												},
												"sslv3_alert_peer_error_no_cert": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert peer error no cert",
												},
												"sslv3_alert_peer_error_no_cipher": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert peer error no cipher",
												},
												"sslv3_alert_peer_error_unsupp_cert_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert peer error unsupp cert type",
												},
												"sslv3_alert_unexpected_msg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert unexpected msg",
												},
												"sslv3_alert_unknown_remote_err_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert unknown remote err type",
												},
												"sslv3_alert_unspported_cert": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for sslv3 alert unspported cert",
												},
												"ssl_ctx_has_no_default_ssl_version": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl ctx has no default ssl version",
												},
												"ssl_handshake_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl handshake failure",
												},
												"ssl_library_has_no_ciphers": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl library has no ciphers",
												},
												"ssl_session_id_callback_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl session id callback failed",
												},
												"ssl_session_id_conflict": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl session id conflict",
												},
												"ssl_session_id_context_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl session id context too long",
												},
												"ssl_session_id_has_bad_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl session id has bad length",
												},
												"ssl_session_id_is_different": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl session id is different",
												},
												"tlsv1_alert_access_denied": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert access denied",
												},
												"tlsv1_alert_decode_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert decode error",
												},
												"tlsv1_alert_decryption_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert decryption failed",
												},
												"tlsv1_alert_decrypt_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert decrypt error",
												},
												"tlsv1_alert_export_restriction": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert export restriction",
												},
												"tlsv1_alert_insufficient_security": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert insufficient security",
												},
												"tlsv1_alert_internal_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert internal error",
												},
												"tlsv1_alert_no_renegotiation": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert no renegotiation",
												},
												"tlsv1_alert_protocol_version": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert protocol version",
												},
												"tlsv1_alert_record_overflow": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert record overflow",
												},
												"tlsv1_alert_unknown_ca": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert unknown ca",
												},
												"tlsv1_alert_user_cancelled": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tlsv1 alert user cancelled",
												},
												"tls_client_cert_req_with_anon_cipher": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tls client cert req with anon cipher",
												},
												"tls_peer_did_not_respond_with_cert_list": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tls peer did not respond with cert list",
												},
												"tls_rsa_encrypted_value_length_is_wrong": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tls rsa encrypted value length is wrong",
												},
												"tried_to_use_unsupported_cipher": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tried to use unsupported cipher",
												},
												"unable_to_decode_dh_certs": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to decode dh certs",
												},
												"unable_to_extract_public_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to extract public key",
												},
												"unable_to_find_dh_parameters": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to find dh parameters",
												},
												"unable_to_find_public_key_parameters": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to find public key parameters",
												},
												"unable_to_find_ssl_method": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to find ssl method",
												},
												"unable_to_load_ssl2_md5_routines": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to load ssl2 md5 routines",
												},
												"unable_to_load_ssl3_md5_routines": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to load ssl3 md5 routines",
												},
												"unable_to_load_ssl3_sha1_routines": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unable to load ssl3 sha1 routines",
												},
												"unexpected_message": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unexpected message",
												},
												"unexpected_record": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unexpected record",
												},
												"uninitialized": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for uninitialized",
												},
												"unknown_alert_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown alert type",
												},
												"unknown_certificate_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown certificate type",
												},
												"unknown_cipher_returned": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown cipher returned",
												},
												"unknown_cipher_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown cipher type",
												},
												"unknown_key_exchange_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown key exchange type",
												},
												"unknown_pkey_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown pkey type",
												},
												"unknown_protocol": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown protocol",
												},
												"unknown_remote_error_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown remote error type",
												},
												"unknown_ssl_version": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown ssl version",
												},
												"unknown_state": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unknown state",
												},
												"unsupported_cipher": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported cipher",
												},
												"unsupported_compression_algorithm": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported compression algorithm",
												},
												"unsupported_option": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported option",
												},
												"unsupported_protocol": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported protocol",
												},
												"unsupported_ssl_version": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported ssl version",
												},
												"unsupported_status_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported status type",
												},
												"write_bio_not_set": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for write bio not set",
												},
												"wrong_cipher_returned": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong cipher returned",
												},
												"wrong_message_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong message type",
												},
												"wrong_counter_of_key_bits": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong counter of key bits",
												},
												"wrong_signature_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong signature length",
												},
												"wrong_signature_size": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong signature size",
												},
												"wrong_ssl_version": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong ssl version",
												},
												"wrong_version_counter": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for wrong version counter",
												},
												"x509_lib": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for x509 lib",
												},
												"x509_verification_setup_problems": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for x509 verification setup problems",
												},
												"clienthello_tlsext": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for clienthello tlsext",
												},
												"parse_tlsext": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for parse tlsext",
												},
												"serverhello_tlsext": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for serverhello tlsext",
												},
												"ssl3_ext_invalid_servername": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl3 ext invalid servername",
												},
												"ssl3_ext_invalid_servername_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ssl3 ext invalid servername type",
												},
												"multiple_sgc_restarts": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for multiple sgc restarts",
												},
												"tls_invalid_ecpointformat_list": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tls invalid ecpointformat list",
												},
												"bad_ecc_cert": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad ecc cert",
												},
												"bad_ecdsa_sig": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad ecdsa sig",
												},
												"bad_ecpoint": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad ecpoint",
												},
												"cookie_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cookie mismatch",
												},
												"unsupported_elliptic_curve": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported elliptic curve",
												},
												"no_required_digest": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no required digest",
												},
												"unsupported_digest_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unsupported digest type",
												},
												"bad_handshake_length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad handshake length",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_ssl_cert_revoke": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ocsp_chain_status_revoked": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Certificate chain status revoked",
												},
												"ocsp_chain_status_unknown": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Certificate chain status unknown",
												},
												"ocsp_connection_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP connection error",
												},
												"ocsp_uri_not_found": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP URI not found",
												},
												"ocsp_uri_https": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP URI https",
												},
												"ocsp_uri_unsupported": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP URI unsupported",
												},
												"ocsp_response_status_revoked": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP response status revoked",
												},
												"ocsp_response_status_unknown": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP response status unknown",
												},
												"ocsp_cache_status_revoked": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP cache status revoked",
												},
												"ocsp_cache_miss": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP cache miss",
												},
												"ocsp_other_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP other errors",
												},
												"ocsp_response_no_nonce": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP other errors",
												},
												"ocsp_response_nonce_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP other errors",
												},
												"crl_connection_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL connection errors",
												},
												"crl_uri_not_found": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL URI not found",
												},
												"crl_uri_https": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL URI https",
												},
												"crl_uri_unsupported": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL URI unsupported",
												},
												"crl_response_status_revoked": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL response status revoked",
												},
												"crl_response_status_unknown": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL response status unknown",
												},
												"crl_cache_status_revoked": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL cache status revoked",
												},
												"crl_other_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL other errors",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"ocsp_chain_status_revoked": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Certificate chain status revoked",
												},
												"ocsp_chain_status_unknown": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Certificate chain status unknown",
												},
												"ocsp_connection_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP connection error",
												},
												"ocsp_uri_not_found": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP URI not found",
												},
												"ocsp_uri_https": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP URI https",
												},
												"ocsp_uri_unsupported": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP URI unsupported",
												},
												"ocsp_response_status_revoked": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP response status revoked",
												},
												"ocsp_response_status_unknown": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP response status unknown",
												},
												"ocsp_cache_status_revoked": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP cache status revoked",
												},
												"ocsp_cache_miss": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for OCSP cache miss",
												},
												"ocsp_other_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP other errors",
												},
												"ocsp_response_no_nonce": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP other errors",
												},
												"ocsp_response_nonce_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log OCSP other errors",
												},
												"crl_connection_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL connection errors",
												},
												"crl_uri_not_found": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL URI not found",
												},
												"crl_uri_https": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL URI https",
												},
												"crl_uri_unsupported": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL URI unsupported",
												},
												"crl_response_status_revoked": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL response status revoked",
												},
												"crl_response_status_unknown": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL response status unknown",
												},
												"crl_cache_status_revoked": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL cache status revoked",
												},
												"crl_other_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CRL other errors",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_ssl_forward_proxy": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"failed_in_ssl_handshakes": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Failed in SSL handshakes",
												},
												"failed_in_crypto_operations": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Failed in crypto operations",
												},
												"failed_in_tcp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Failed in TCP",
												},
												"failed_in_certificate_verification": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Failed in Certificate verification",
												},
												"failed_in_certificate_signing": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Failed in Certificate signing",
												},
												"invalid_ocsp_stapling_response": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid OCSP Stapling Response",
												},
												"revoked_ocsp_response": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Revoked OCSP Response",
												},
												"unsupported_ssl_version": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported SSL version",
												},
												"connections_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Connections failed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"failed_in_ssl_handshakes": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Failed in SSL handshakes",
												},
												"failed_in_crypto_operations": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Failed in crypto operations",
												},
												"failed_in_tcp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Failed in TCP",
												},
												"failed_in_certificate_verification": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Failed in Certificate verification",
												},
												"failed_in_certificate_signing": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Failed in Certificate signing",
												},
												"invalid_ocsp_stapling_response": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid OCSP Stapling Response",
												},
												"revoked_ocsp_response": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Revoked OCSP Response",
												},
												"unsupported_ssl_version": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported SSL version",
												},
												"connections_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Connections failed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"vpn_error": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"bad_opcode": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_opcode",
												},
												"bad_sg_write_len": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_sg_write_len",
												},
												"bad_len": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_len",
												},
												"bad_ipsec_protocol": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_protocol",
												},
												"bad_ipsec_auth": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_auth",
												},
												"bad_ipsec_padding": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_padding",
												},
												"bad_ip_version": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ip_version",
												},
												"bad_auth_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_auth_type",
												},
												"bad_encrypt_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_encrypt_type",
												},
												"bad_ipsec_spi": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_spi",
												},
												"bad_checksum": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_checksum",
												},
												"bad_ipsec_context": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_context",
												},
												"bad_ipsec_context_direction": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_context_direction",
												},
												"bad_ipsec_context_flag_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_context_flag_mismatch",
												},
												"ipcomp_payload": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipcomp_payload",
												},
												"bad_selector_match": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_selector_match",
												},
												"bad_fragment_size": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_fragment_size",
												},
												"bad_inline_data": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_inline_data",
												},
												"bad_frag_size_configuration": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_frag_size_configuration",
												},
												"dummy_payload": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for dummy_payload",
												},
												"bad_ip_payload_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ip_payload_type",
												},
												"bad_min_frag_size_auth_sha384_512": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_min_frag_size_auth_sha384_512",
												},
												"bad_esp_next_header": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_esp_next_header",
												},
												"bad_gre_header": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_gre_header",
												},
												"bad_gre_protocol": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_gre_protocol",
												},
												"ipv6_extension_headers_too_big": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_extension_headers_too_big",
												},
												"ipv6_hop_by_hop_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_hop_by_hop_error",
												},
												"error_ipv6_decrypt_rh_segs_left_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for error_ipv6_decrypt_rh_segs_left_error",
												},
												"ipv6_rh_length_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_rh_length_error",
												},
												"ipv6_outbound_rh_copy_addr_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_outbound_rh_copy_addr_error",
												},
												"error_ipv6_extension_header_bad": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for error_IPv6_extension_header_bad",
												},
												"bad_encrypt_type_ctr_gcm": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_encrypt_type_ctr_gcm",
												},
												"ah_not_supported_with_gcm_gmac_sha2": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ah_not_supported_with_gcm_gmac_sha2",
												},
												"tfc_padding_with_prefrag_not_supported": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tfc_padding_with_prefrag_not_supported",
												},
												"bad_srtp_auth_tag": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_srtp_auth_tag",
												},
												"bad_ipcomp_configuration": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipcomp_configuration",
												},
												"dsiv_incorrect_param": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for dsiv_incorrect_param",
												},
												"bad_ipsec_unknown": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_unknown",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"bad_opcode": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_opcode",
												},
												"bad_sg_write_len": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_sg_write_len",
												},
												"bad_len": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_len",
												},
												"bad_ipsec_protocol": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_protocol",
												},
												"bad_ipsec_auth": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_auth",
												},
												"bad_ipsec_padding": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_padding",
												},
												"bad_ip_version": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ip_version",
												},
												"bad_auth_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_auth_type",
												},
												"bad_encrypt_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_encrypt_type",
												},
												"bad_ipsec_spi": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_spi",
												},
												"bad_checksum": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_checksum",
												},
												"bad_ipsec_context": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_context",
												},
												"bad_ipsec_context_direction": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_context_direction",
												},
												"bad_ipsec_context_flag_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_context_flag_mismatch",
												},
												"ipcomp_payload": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipcomp_payload",
												},
												"bad_selector_match": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_selector_match",
												},
												"bad_fragment_size": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_fragment_size",
												},
												"bad_inline_data": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_inline_data",
												},
												"bad_frag_size_configuration": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_frag_size_configuration",
												},
												"dummy_payload": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for dummy_payload",
												},
												"bad_ip_payload_type": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ip_payload_type",
												},
												"bad_min_frag_size_auth_sha384_512": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_min_frag_size_auth_sha384_512",
												},
												"bad_esp_next_header": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_esp_next_header",
												},
												"bad_gre_header": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_gre_header",
												},
												"bad_gre_protocol": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_gre_protocol",
												},
												"ipv6_extension_headers_too_big": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_extension_headers_too_big",
												},
												"ipv6_hop_by_hop_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_hop_by_hop_error",
												},
												"error_ipv6_decrypt_rh_segs_left_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for error_ipv6_decrypt_rh_segs_left_error",
												},
												"ipv6_rh_length_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_rh_length_error",
												},
												"ipv6_outbound_rh_copy_addr_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ipv6_outbound_rh_copy_addr_error",
												},
												"error_ipv6_extension_header_bad": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for error_IPv6_extension_header_bad",
												},
												"bad_encrypt_type_ctr_gcm": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_encrypt_type_ctr_gcm",
												},
												"ah_not_supported_with_gcm_gmac_sha2": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ah_not_supported_with_gcm_gmac_sha2",
												},
												"tfc_padding_with_prefrag_not_supported": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tfc_padding_with_prefrag_not_supported",
												},
												"bad_srtp_auth_tag": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_srtp_auth_tag",
												},
												"bad_ipcomp_configuration": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipcomp_configuration",
												},
												"dsiv_incorrect_param": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for dsiv_incorrect_param",
												},
												"bad_ipsec_unknown": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for bad_ipsec_unknown",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_ddos_proc": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"l3_entry_match_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Entry match drop",
												},
												"l3_entry_match_drop_hw": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 HW entry match drop",
												},
												"l3_entry_drop_max_hw_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Entry Drop due to HW Limit Exceeded",
												},
												"l4_entry_match_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 Entry match drop",
												},
												"l4_entry_match_drop_hw": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 HW Entry match drop",
												},
												"l4_entry_drop_max_hw_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 Entry Drop due to HW Limit Exceeded",
												},
												"l4_entry_list_alloc_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 Entry list alloc failures",
												},
												"ip_node_alloc_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Node alloc failures",
												},
												"ip_port_block_alloc_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port block alloc failure",
												},
												"ip_other_block_alloc_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Other block alloc failure",
												},
												"l3_entry_add_to_bgp_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Entry BGP add failures",
												},
												"l3_entry_remove_from_bgp_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 entry BGP remove failures",
												},
												"l3_entry_add_to_hw_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 entry HW add failure",
												},
												"syn_cookie_verification_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SYN cookie verification failed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"l3_entry_match_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Entry match drop",
												},
												"l3_entry_match_drop_hw": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 HW entry match drop",
												},
												"l3_entry_drop_max_hw_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Entry Drop due to HW Limit Exceeded",
												},
												"l4_entry_match_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 Entry match drop",
												},
												"l4_entry_match_drop_hw": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 HW Entry match drop",
												},
												"l4_entry_drop_max_hw_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 Entry Drop due to HW Limit Exceeded",
												},
												"l4_entry_list_alloc_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 Entry list alloc failures",
												},
												"ip_node_alloc_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Node alloc failures",
												},
												"ip_port_block_alloc_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port block alloc failure",
												},
												"ip_other_block_alloc_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Other block alloc failure",
												},
												"l3_entry_add_to_bgp_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Entry BGP add failures",
												},
												"l3_entry_remove_from_bgp_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 entry BGP remove failures",
												},
												"l3_entry_add_to_hw_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 entry HW add failure",
												},
												"syn_cookie_verification_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SYN cookie verification failed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_lsn": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"user_quota_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Creation Failed",
												},
												"data_sesn_user_quota_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Session User-Quota Exceeded",
												},
												"fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Full-cone Session Creation Failed",
												},
												"fullcone_self_hairpinning_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Self-Hairpinning Drop",
												},
												"nat_pool_unusable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT Pool Unusable",
												},
												"ha_nat_pool_unusable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA NAT Pool Unusable",
												},
												"ha_nat_pool_batch_type_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA NAT Pool Batch Type Mismatch",
												},
												"sip_alg_quota_inc_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG User-Quota Exceeded",
												},
												"sip_alg_alloc_rtp_rtcp_port_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Alloc RTP/RTCP NAT Ports Failure",
												},
												"sip_alg_alloc_single_port_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Alloc Single RTP or RTCP NAT Port Failure",
												},
												"sip_alg_create_single_fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Create Single RTP or RTCP Full-cone Session Failure",
												},
												"sip_alg_create_rtp_fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Create RTP Full-cone Session Failure",
												},
												"sip_alg_create_rtcp_fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Create RTCP Full-cone Session Failure",
												},
												"h323_alg_alloc_single_port_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Alloc Single RTP or RTCP NAT Port Failure",
												},
												"h323_alg_create_single_fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Create Single RTP or RTCP Full-cone Session Failure",
												},
												"h323_alg_create_rtp_fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Create RTP Full-cone Session Failure",
												},
												"h323_alg_create_rtcp_fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Create RTCP Full-cone Session Failure",
												},
												"port_overloading_out_of_memory": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Overloading Out of Memory",
												},
												"port_overloading_inc_overflow": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Overloading Inc Overflow",
												},
												"fullcone_ext_mem_alloc_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for LSN Fullcone Extension Memory Allocate Failure",
												},
												"fullcone_ext_mem_alloc_init_faulure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for LSN Fullcone Extension Initialization Failure",
												},
												"mgcp_alg_create_rtp_fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for MGCP ALG Create RTP Full-cone Session Failure",
												},
												"mgcp_alg_create_rtcp_fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for MGCP ALG Create RTCP Full-cone Session Failure",
												},
												"mgcp_alg_port_pair_alloc_from_quota_par": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for MGCP ALG Port Pair Allocated From Quota Partition Error",
												},
												"user_quota_unusable_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Unusable Drop",
												},
												"user_quota_unusable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Marked Unusable",
												},
												"adc_port_allocation_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ADC Port Allocation Failed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"user_quota_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Creation Failed",
												},
												"data_sesn_user_quota_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Session User-Quota Exceeded",
												},
												"fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Full-cone Session Creation Failed",
												},
												"fullcone_self_hairpinning_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Self-Hairpinning Drop",
												},
												"nat_pool_unusable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT Pool Unusable",
												},
												"ha_nat_pool_unusable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA NAT Pool Unusable",
												},
												"ha_nat_pool_batch_type_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA NAT Pool Batch Type Mismatch",
												},
												"sip_alg_quota_inc_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG User-Quota Exceeded",
												},
												"sip_alg_alloc_rtp_rtcp_port_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Alloc RTP/RTCP NAT Ports Failure",
												},
												"sip_alg_alloc_single_port_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Alloc Single RTP or RTCP NAT Port Failure",
												},
												"sip_alg_create_single_fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Create Single RTP or RTCP Full-cone Session Failure",
												},
												"sip_alg_create_rtp_fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Create RTP Full-cone Session Failure",
												},
												"sip_alg_create_rtcp_fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP ALG Create RTCP Full-cone Session Failure",
												},
												"h323_alg_alloc_single_port_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Alloc Single RTP or RTCP NAT Port Failure",
												},
												"h323_alg_create_single_fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Create Single RTP or RTCP Full-cone Session Failure",
												},
												"h323_alg_create_rtp_fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Create RTP Full-cone Session Failure",
												},
												"h323_alg_create_rtcp_fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for H323 ALG Create RTCP Full-cone Session Failure",
												},
												"port_overloading_out_of_memory": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Overloading Out of Memory",
												},
												"port_overloading_inc_overflow": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port Overloading Inc Overflow",
												},
												"fullcone_ext_mem_alloc_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for LSN Fullcone Extension Memory Allocate Failure",
												},
												"fullcone_ext_mem_alloc_init_faulure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for LSN Fullcone Extension Initialization Failure",
												},
												"mgcp_alg_create_rtp_fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for MGCP ALG Create RTP Full-cone Session Failure",
												},
												"mgcp_alg_create_rtcp_fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for MGCP ALG Create RTCP Full-cone Session Failure",
												},
												"mgcp_alg_port_pair_alloc_from_quota_par": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for MGCP ALG Port Pair Allocated From Quota Partition Error",
												},
												"user_quota_unusable_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Unusable Drop",
												},
												"user_quota_unusable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Marked Unusable",
												},
												"adc_port_allocation_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ADC Port Allocation Failed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_lsn_alg_esp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"nat_ip_conflict": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT IP Conflict",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"nat_ip_conflict": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT IP Conflict",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_lsn_alg_pptp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"no_gre_session_match": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Matching GRE Session",
												},
												"call_req_pns_call_id_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Call ID Mismatch on Call Request",
												},
												"call_reply_pns_call_id_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Call ID Mismatch on Call Reply",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"no_gre_session_match": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Matching GRE Session",
												},
												"call_req_pns_call_id_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Call ID Mismatch on Call Request",
												},
												"call_reply_pns_call_id_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Call ID Mismatch on Call Reply",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_lsn_alg_rtsp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"stream_creation_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Stream Creation Failures",
												},
												"port_allocation_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Stream Client Port Allocation Failures",
												},
												"unknown_client_port_from_server": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server Replies With Unknown Client Ports",
												},
												"no_session_mem": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Session Creation Failures",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"stream_creation_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Stream Creation Failures",
												},
												"port_allocation_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Stream Client Port Allocation Failures",
												},
												"unknown_client_port_from_server": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server Replies With Unknown Client Ports",
												},
												"no_session_mem": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Session Creation Failures",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_lsn_alg_sip": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"method_unknown": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP Method UNKNOWN",
												},
												"parse_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP Message Parse Error",
												},
												"tcp_out_of_order_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Out-of-Order Drop",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"method_unknown": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP Method UNKNOWN",
												},
												"parse_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP Message Parse Error",
												},
												"tcp_out_of_order_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Out-of-Order Drop",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_lsn_alg_mgcp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parse_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for MGCP Message Parse Error",
												},
												"tcp_out_of_order_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Out-of-Order Drop",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"parse_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for MGCP Message Parse Error",
												},
												"tcp_out_of_order_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Out-of-Order Drop",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_lsn_alg_h323": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parse_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Message Parse Error",
												},
												"tcp_out_of_order_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Out-of-Order Drop",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"parse_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Message Parse Error",
												},
												"tcp_out_of_order_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Out-of-Order Drop",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_lsn_radius": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"radius_request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Dropped (Malformed Packet)",
												},
												"request_bad_secret_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Bad Secret Dropped",
												},
												"request_no_key_vap_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request No Key Attribute Dropped",
												},
												"request_malformed_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Malformed Dropped",
												},
												"request_ignored": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Ignored",
												},
												"radius_table_full": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Dropped (Table Full)",
												},
												"secret_not_configured_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Secret Not Configured Dropped",
												},
												"ha_standby_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA Standby Dropped",
												},
												"invalid_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Radius Request has Invalid Key Field",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"radius_request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Dropped (Malformed Packet)",
												},
												"request_bad_secret_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Bad Secret Dropped",
												},
												"request_no_key_vap_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request No Key Attribute Dropped",
												},
												"request_malformed_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Malformed Dropped",
												},
												"request_ignored": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Ignored",
												},
												"radius_table_full": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Dropped (Table Full)",
												},
												"secret_not_configured_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Secret Not Configured Dropped",
												},
												"ha_standby_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA Standby Dropped",
												},
												"invalid_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Radius Request has Invalid Key Field",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_nat64_global": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"user_quota_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Creation Failed",
												},
												"nat_port_unavailable_tcp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP NAT Port Unavailable",
												},
												"nat_port_unavailable_udp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP NAT Port Unavailable",
												},
												"nat_port_unavailable_icmp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMP NAT Port Unavailable",
												},
												"new_user_resource_unavailable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for New User NAT Resource Unavailable",
												},
												"fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Full-cone Session Creation Failed",
												},
												"fullcone_self_hairpinning_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Self-Hairpinning Drop",
												},
												"eif_limit_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Endpoint-Independent Filtering Inbound Limit Exceeded",
												},
												"nat_pool_unusable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT Pool Unusable",
												},
												"ha_nat_pool_unusable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA NAT Pool Unusable",
												},
												"ha_nat_pool_batch_type_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA NAT Pool Batch Type Mismatch",
												},
												"no_radius_profile_match": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No RADIUS Profile Match",
												},
												"no_class_list_match": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Class-List Match",
												},
												"user_quota_unusable_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Unusable Drop",
												},
												"user_quota_unusable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Marked Unusable",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"user_quota_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Creation Failed",
												},
												"nat_port_unavailable_tcp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP NAT Port Unavailable",
												},
												"nat_port_unavailable_udp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP NAT Port Unavailable",
												},
												"nat_port_unavailable_icmp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMP NAT Port Unavailable",
												},
												"new_user_resource_unavailable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for New User NAT Resource Unavailable",
												},
												"fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Full-cone Session Creation Failed",
												},
												"fullcone_self_hairpinning_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Self-Hairpinning Drop",
												},
												"eif_limit_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Endpoint-Independent Filtering Inbound Limit Exceeded",
												},
												"nat_pool_unusable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT Pool Unusable",
												},
												"ha_nat_pool_unusable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA NAT Pool Unusable",
												},
												"ha_nat_pool_batch_type_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA NAT Pool Batch Type Mismatch",
												},
												"no_radius_profile_match": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No RADIUS Profile Match",
												},
												"no_class_list_match": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Class-List Match",
												},
												"user_quota_unusable_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Unusable Drop",
												},
												"user_quota_unusable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Marked Unusable",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_ds_lite_global": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"user_quota_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Creation Failed",
												},
												"nat_port_unavailable_tcp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP NAT Port Unavailable",
												},
												"nat_port_unavailable_udp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP NAT Port Unavailable",
												},
												"nat_port_unavailable_icmp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMP NAT Port Unavailable",
												},
												"fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Full-cone Session Creation Failed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"user_quota_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for User-Quota Creation Failed",
												},
												"nat_port_unavailable_tcp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP NAT Port Unavailable",
												},
												"nat_port_unavailable_udp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP NAT Port Unavailable",
												},
												"nat_port_unavailable_icmp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMP NAT Port Unavailable",
												},
												"fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Full-cone Session Creation Failed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_fixed_nat_global": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"nat_port_unavailable_tcp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP NAT Port Unavailable",
												},
												"nat_port_unavailable_udp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP NAT Port Unavailable",
												},
												"nat_port_unavailable_icmp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMP NAT Port Unavailable",
												},
												"session_user_quota_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Sessions User Quota Exceeded",
												},
												"fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Full-Cone Session Creation Failed",
												},
												"nat44_inbound_filtered": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT44 Endpoint-Dependent Filtering Drop",
												},
												"nat64_inbound_filtered": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT64 Endpoint-Dependent Filtering Drop",
												},
												"dslite_inbound_filtered": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DS-Lite Endpoint-Dependent Filtering Drop",
												},
												"nat44_eif_limit_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT44 Endpoint-Independent-Filtering Limit Exceeded",
												},
												"nat64_eif_limit_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT64 Endpoint-Independent-Filtering Limit Exceeded",
												},
												"dslite_eif_limit_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DS-Lite Endpoint-Independent-Filtering Limit Exceeded",
												},
												"standby_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT LID Standby Drop",
												},
												"fixed_nat_fullcone_self_hairpinning_dro": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Self-Hairpinning Drop",
												},
												"sixrd_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT IPv6 in IPv4 Packet Drop",
												},
												"dest_rlist_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT Dest Rule List Drop",
												},
												"dest_rlist_pass_through": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT Dest Rule List Pass-Through",
												},
												"dest_rlist_snat_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT Dest Rules List Source NAT Drop",
												},
												"config_not_found": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT Config not Found",
												},
												"port_overload_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port overload failed",
												},
												"ha_session_user_quota_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA Sessions User Quota Exceeded",
												},
												"unused_ip_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT UNUSED IP Hit and Drop",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"nat_port_unavailable_tcp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP NAT Port Unavailable",
												},
												"nat_port_unavailable_udp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for UDP NAT Port Unavailable",
												},
												"nat_port_unavailable_icmp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMP NAT Port Unavailable",
												},
												"session_user_quota_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Sessions User Quota Exceeded",
												},
												"fullcone_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Full-Cone Session Creation Failed",
												},
												"nat44_inbound_filtered": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT44 Endpoint-Dependent Filtering Drop",
												},
												"nat64_inbound_filtered": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT64 Endpoint-Dependent Filtering Drop",
												},
												"dslite_inbound_filtered": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DS-Lite Endpoint-Dependent Filtering Drop",
												},
												"nat44_eif_limit_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT44 Endpoint-Independent-Filtering Limit Exceeded",
												},
												"nat64_eif_limit_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT64 Endpoint-Independent-Filtering Limit Exceeded",
												},
												"dslite_eif_limit_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DS-Lite Endpoint-Independent-Filtering Limit Exceeded",
												},
												"standby_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT LID Standby Drop",
												},
												"fixed_nat_fullcone_self_hairpinning_dro": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Self-Hairpinning Drop",
												},
												"sixrd_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT IPv6 in IPv4 Packet Drop",
												},
												"dest_rlist_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT Dest Rule List Drop",
												},
												"dest_rlist_pass_through": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT Dest Rule List Pass-Through",
												},
												"dest_rlist_snat_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT Dest Rules List Source NAT Drop",
												},
												"config_not_found": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT Config not Found",
												},
												"port_overload_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Port overload failed",
												},
												"ha_session_user_quota_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA Sessions User Quota Exceeded",
												},
												"unused_ip_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fixed NAT UNUSED IP Hit and Drop",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_fixed_nat_alg_pptp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"call_req_pns_call_id_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Call ID Mismatch on Call Request",
												},
												"call_reply_pns_call_id_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Call ID Mismatch on Call Reply",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"call_req_pns_call_id_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Call ID Mismatch on Call Request",
												},
												"call_reply_pns_call_id_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Call ID Mismatch on Call Reply",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_fixed_nat_alg_rtsp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"stream_creation_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Stream Creation Failures",
												},
												"port_allocation_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Stream Client Port Allocation Failures",
												},
												"no_session_mem": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Session Creation Failures",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"stream_creation_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Stream Creation Failures",
												},
												"port_allocation_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Stream Client Port Allocation Failures",
												},
												"no_session_mem": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Session Creation Failures",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_fixed_nat_alg_sip": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"method_unknown": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP Method UNKNOWN",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"method_unknown": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SIP Method UNKNOWN",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_pcp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"pkt_not_request_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Not a PCP Request",
												},
												"pkt_too_short_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Too Short",
												},
												"noroute_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response No Route",
												},
												"unsupported_version": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PCP version",
												},
												"not_authorized": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for PCP Request Not Authorized",
												},
												"malform_request": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for PCP Request Malformed",
												},
												"unsupp_opcode": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PCP Opcode",
												},
												"unsupp_option": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PCP Option",
												},
												"malform_option": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for PCP Option Malformed",
												},
												"no_resources": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No System or NAT Resources",
												},
												"unsupp_protocol": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported Mapping Protocol",
												},
												"cannot_provide_suggest": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cannot Provide Suggested Port When PREFER_FAILURE",
												},
												"address_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for PCP Client Address Mismatch",
												},
												"excessive_remote_peers": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Excessive Remote Peers",
												},
												"pkt_not_from_nat_inside": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Dropped For Not Coming From NAT Inside",
												},
												"l4_process_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3/L4 Process Error",
												},
												"internal_error_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Internal Error",
												},
												"unsol_ance_sent_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsolicited Announce Send Failure",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"pkt_not_request_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Not a PCP Request",
												},
												"pkt_too_short_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Too Short",
												},
												"noroute_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response No Route",
												},
												"unsupported_version": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PCP version",
												},
												"not_authorized": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for PCP Request Not Authorized",
												},
												"malform_request": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for PCP Request Malformed",
												},
												"unsupp_opcode": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PCP Opcode",
												},
												"unsupp_option": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PCP Option",
												},
												"malform_option": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for PCP Option Malformed",
												},
												"no_resources": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No System or NAT Resources",
												},
												"unsupp_protocol": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported Mapping Protocol",
												},
												"cannot_provide_suggest": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cannot Provide Suggested Port When PREFER_FAILURE",
												},
												"address_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for PCP Client Address Mismatch",
												},
												"excessive_remote_peers": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Excessive Remote Peers",
												},
												"pkt_not_from_nat_inside": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Dropped For Not Coming From NAT Inside",
												},
												"l4_process_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3/L4 Process Error",
												},
												"internal_error_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Internal Error",
												},
												"unsol_ance_sent_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsolicited Announce Send Failure",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_logging": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"log_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log Packets Dropped",
												},
												"conn_tcp_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Connection Lost",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"log_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log Packets Dropped",
												},
												"conn_tcp_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP Connection Lost",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_l4": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"out_of_session_memory": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of Session Memory",
												},
												"icmp_host_unreachable_sent": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMP Host Unreachable Sent",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"out_of_session_memory": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of Session Memory",
												},
												"icmp_host_unreachable_sent": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMP Host Unreachable Sent",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_icmp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"icmp_to_icmp_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMP to ICMP Conversion Error",
												},
												"icmp_to_icmpv6_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMP to ICMPv6 Conversion Error",
												},
												"icmpv6_to_icmp_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMPv6 to ICMP Conversion Error",
												},
												"icmpv6_to_icmpv6_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMPv6 to ICMPv6 Conversion Error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"icmp_to_icmp_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMP to ICMP Conversion Error",
												},
												"icmp_to_icmpv6_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMP to ICMPv6 Conversion Error",
												},
												"icmpv6_to_icmp_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMPv6 to ICMP Conversion Error",
												},
												"icmpv6_to_icmpv6_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICMPv6 to ICMPv6 Conversion Error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_http_alg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"radius_requst_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Query Request Dropped",
												},
												"radius_response_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Query Response Dropped",
												},
												"out_of_memory_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out-of-Memory Dropped",
												},
												"queue_len_exceed_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Queue Length Exceed Dropped",
												},
												"out_of_order_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Out-of-Order Dropped",
												},
												"header_insertion_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Buff Insertion Failed",
												},
												"header_removal_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Buff Removal Failed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"radius_requst_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Query Request Dropped",
												},
												"radius_response_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Query Response Dropped",
												},
												"out_of_memory_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out-of-Memory Dropped",
												},
												"queue_len_exceed_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Queue Length Exceed Dropped",
												},
												"out_of_order_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packet Out-of-Order Dropped",
												},
												"header_insertion_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Buff Insertion Failed",
												},
												"header_removal_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Buff Removal Failed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_dns64": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"query_bad_pkt": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Query Bad Packet",
												},
												"resp_bad_pkt": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response Bad Packet",
												},
												"resp_bad_qr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response Bad Query",
												},
												"drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Dropped",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"query_bad_pkt": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Query Bad Packet",
												},
												"resp_bad_pkt": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response Bad Packet",
												},
												"resp_bad_qr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response Bad Query",
												},
												"drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Dropped",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"cgnv6_dhcpv6": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"packets_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packets Dropped",
												},
												"pkts_dropped_during_clear": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packets dropped during clear bindings process",
												},
												"rcv_not_supported_msg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packets with not supported messages sent",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"packets_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packets Dropped",
												},
												"pkts_dropped_during_clear": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packets dropped during clear bindings process",
												},
												"rcv_not_supported_msg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Packets with not supported messages sent",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"fw_logging": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"log_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log Packets Dropped",
												},
												"http_logging_invalid_format": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP Logging Invalid Format Error",
												},
												"session_limit_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Session Limit Exceeded",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"log_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Log Packets Dropped",
												},
												"http_logging_invalid_format": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP Logging Invalid Format Error",
												},
												"session_limit_exceeded": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Session Limit Exceeded",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"fw_global": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"fullcone_creation_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Full-Cone Creation Failure",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"fullcone_creation_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Full-Cone Creation Failure",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"fw_alg_rtsp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"transport_alloc_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Transport Alloc Failure",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"transport_alloc_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Transport Alloc Failure",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"fw_alg_pptp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"call_req_pns_call_id_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Call ID Mismatch on Call Request",
												},
												"call_reply_pns_call_id_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Call ID Mismatch on Call Reply",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"call_req_pns_call_id_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Call ID Mismatch on Call Request",
												},
												"call_reply_pns_call_id_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Call ID Mismatch on Call Reply",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"fw_rad_server": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"radius_request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Dropped (Malformed Packet)",
												},
												"request_bad_secret_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Bad Secret Dropped",
												},
												"request_no_key_vap_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request No Key Attribute Dropped",
												},
												"request_malformed_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Malformed Dropped",
												},
												"request_ignored": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Table Full Dropped",
												},
												"radius_table_full": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Dropped (Table Full)",
												},
												"ha_standby_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA Standby Dropped",
												},
												"ipv6_prefix_length_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Framed IPV6 Prefix Length Mismatch",
												},
												"invalid_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Radius Request has Invalid Key Field",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"radius_request_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Dropped (Malformed Packet)",
												},
												"request_bad_secret_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Bad Secret Dropped",
												},
												"request_no_key_vap_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request No Key Attribute Dropped",
												},
												"request_malformed_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Malformed Dropped",
												},
												"request_ignored": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Table Full Dropped",
												},
												"radius_table_full": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for RADIUS Request Dropped (Table Full)",
												},
												"ha_standby_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HA Standby Dropped",
												},
												"ipv6_prefix_length_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Framed IPV6 Prefix Length Mismatch",
												},
												"invalid_key": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Radius Request has Invalid Key Field",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"fw_tcp_syn_cookie": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"verification_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SYN cookie verification failed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"verification_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SYN cookie verification failed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"fw_ddos_protection": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ddos_entries_too_many": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Too many DDOS entries",
												},
												"ddos_entry_add_to_bgp_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DDoS Entry BGP add failures",
												},
												"ddos_entry_remove_from_bgp_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DDOS entry BGP remove failures",
												},
												"ddos_packet_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DDOS Packet Drop",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"ddos_entries_too_many": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Too many DDOS entries",
												},
												"ddos_entry_add_to_bgp_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DDoS Entry BGP add failures",
												},
												"ddos_entry_remove_from_bgp_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DDOS entry BGP remove failures",
												},
												"ddos_packet_dropped": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DDOS Packet Drop",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"fw_gtp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"out_of_session_memory": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of Tunnel Memory for GTP-C",
												},
												"blade_out_of_session_memory": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of Tunnel Memory on PU2",
												},
												"gtp_smp_path_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP PATH check Failed",
												},
												"gtp_smp_c_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP check Failed",
												},
												"blade_gtp_smp_path_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP PATH check Failed on PU2",
												},
												"blade_gtp_smp_c_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP check Failed on PU2",
												},
												"gtp_tunnel_rate_limit_entry_create_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Tunnel Level Rate Limit Entry Create Failure",
												},
												"gtp_u_tunnel_rate_limit_entry_create_fa": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U Tunnel Level Rate Limit Entry Create Failure",
												},
												"gtp_rate_limit_smp_create_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit SMP Create Failure",
												},
												"gtp_rate_limit_t3_ctr_create_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Dynamic Counters Create Failure",
												},
												"gtp_rate_limit_entry_create_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Entry Create Failure",
												},
												"blade_gtp_rate_limit_smp_create_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit SMP Create Failure on PU2",
												},
												"blade_gtp_rate_limit_t3_ctr_create_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Dynamic Counters Create Failure on PU2",
												},
												"blade_gtp_rate_limit_entry_create_failu": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Entry Create Failure on PU2",
												},
												"gtp_smp_dec_sess_count_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U session count is 0 in GTP-C SMP",
												},
												"gtp_u_smp_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP U-SMP check Failed",
												},
												"gtp_info_ext_not_found": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-Info ext not found while freeing C-smp",
												},
												"blade_gtp_smp_dec_sess_count_check_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U session count is 0 in GTP-C SMP on PU2",
												},
												"blade_gtp_u_smp_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP U-SMP check Failed on PU2",
												},
												"blade_gtp_info_ext_not_found": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-Info ext not found while freeing C-smp on PU2",
												},
												"blade_gtp_smp_session_count_check_faile": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U session count is not in range of 0-11 in GTP-C SMP on PU2",
												},
												"gtp_c_smp_sig_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP signature check Failed",
												},
												"blade_gtp_c_smp_sig_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP signature check Failed on PU2",
												},
												"gtp_u_smp_sig_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP signature check Failed",
												},
												"blade_gtp_u_smp_sig_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U SMP signature check Failed on PU2",
												},
												"gtp_smp_sig_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP signature check Failed",
												},
												"blade_gtp_smp_sig_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP signature check Failed on PU2",
												},
												"gtp_c_fail_conn_create_slow": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C packet failed creating L4-session in slowpath",
												},
												"gtp_u_fail_conn_create_slow": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U packet failed while creating L4-session in slowpath",
												},
												"gtp_pathm_fail_conn_create_slow": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP path packet failed while creating L4-session in slowpath",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"out_of_session_memory": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of Tunnel Memory for GTP-C",
												},
												"blade_out_of_session_memory": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of Tunnel Memory on PU2",
												},
												"gtp_smp_path_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP PATH check Failed",
												},
												"gtp_smp_c_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP check Failed",
												},
												"blade_gtp_smp_path_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP PATH check Failed on PU2",
												},
												"blade_gtp_smp_c_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP check Failed on PU2",
												},
												"gtp_tunnel_rate_limit_entry_create_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Tunnel Level Rate Limit Entry Create Failure",
												},
												"gtp_u_tunnel_rate_limit_entry_create_fa": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U Tunnel Level Rate Limit Entry Create Failure",
												},
												"gtp_rate_limit_smp_create_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit SMP Create Failure",
												},
												"gtp_rate_limit_t3_ctr_create_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Dynamic Counters Create Failure",
												},
												"gtp_rate_limit_entry_create_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Entry Create Failure",
												},
												"blade_gtp_rate_limit_smp_create_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit SMP Create Failure on PU2",
												},
												"blade_gtp_rate_limit_t3_ctr_create_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Dynamic Counters Create Failure on PU2",
												},
												"blade_gtp_rate_limit_entry_create_failu": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Entry Create Failure on PU2",
												},
												"gtp_smp_dec_sess_count_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U session count is 0 in GTP-C SMP",
												},
												"gtp_u_smp_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP U-SMP check Failed",
												},
												"gtp_info_ext_not_found": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-Info ext not found while freeing C-smp",
												},
												"blade_gtp_smp_dec_sess_count_check_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U session count is 0 in GTP-C SMP on PU2",
												},
												"blade_gtp_u_smp_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP U-SMP check Failed on PU2",
												},
												"blade_gtp_info_ext_not_found": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-Info ext not found while freeing C-smp on PU2",
												},
												"blade_gtp_smp_session_count_check_faile": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U session count is not in range of 0-11 in GTP-C SMP on PU2",
												},
												"gtp_c_smp_sig_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP signature check Failed",
												},
												"blade_gtp_c_smp_sig_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP signature check Failed on PU2",
												},
												"gtp_u_smp_sig_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP signature check Failed",
												},
												"blade_gtp_u_smp_sig_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U SMP signature check Failed on PU2",
												},
												"gtp_smp_sig_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP signature check Failed",
												},
												"blade_gtp_smp_sig_check_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP signature check Failed on PU2",
												},
												"gtp_c_fail_conn_create_slow": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C packet failed creating L4-session in slowpath",
												},
												"gtp_u_fail_conn_create_slow": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U packet failed while creating L4-session in slowpath",
												},
												"gtp_pathm_fail_conn_create_slow": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP path packet failed while creating L4-session in slowpath",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"system_tcp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"attemptfails": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Connect attemp failures",
												},
												"noroute": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCPIP out noroute",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"attemptfails": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Connect attemp failures",
												},
												"noroute": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCPIP out noroute",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_conn_reuse": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ntermi_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total terminated by err",
												},
												"pause_conn_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Pause request fail",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"ntermi_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total terminated by err",
												},
												"pause_conn_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Pause request fail",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_aflow": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"pause_conn_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Pause connection fail",
												},
												"error_resume_conn": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resume conn by error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"pause_conn_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Pause connection fail",
												},
												"error_resume_conn": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resume conn by error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_fix": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"svrsel_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
												},
												"noroute": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No route failure",
												},
												"snat_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT failure",
												},
												"client_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client fail",
												},
												"server_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server fail",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"svrsel_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
												},
												"noroute": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No route failure",
												},
												"snat_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT failure",
												},
												"client_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client fail",
												},
												"server_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server fail",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_spdy_proxy": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"tcp_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP sock error",
												},
												"stream_not_found": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for STREAM not found",
												},
												"stream_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Stream err",
												},
												"session_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Session err",
												},
												"data_no_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data no stream found",
												},
												"data_no_stream_no_goaway": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data no stream and no goaway",
												},
												"data_no_stream_goaway_close": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data no stream and no goaway and close session",
												},
												"est_cb_no_tuple": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Est callback no tuple",
												},
												"data_cb_no_tuple": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data callback no tuple",
												},
												"ctx_alloc_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Context alloc fail",
												},
												"stream_alloc_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Stream alloc fail",
												},
												"http_conn_alloc_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP connection allocation fail",
												},
												"request_header_alloc_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Request/Header allocation fail",
												},
												"decompress_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Decompress fail",
												},
												"invalid_frame_size": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid frame size",
												},
												"invalid_version": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid version",
												},
												"compress_ctx_alloc_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Compression context allocation fail",
												},
												"header_compress_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Header compress fail",
												},
												"http_err_stream_closed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP error stream already closed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"tcp_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP sock error",
												},
												"stream_not_found": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for STREAM not found",
												},
												"stream_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Stream err",
												},
												"session_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Session err",
												},
												"data_no_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data no stream found",
												},
												"data_no_stream_no_goaway": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data no stream and no goaway",
												},
												"data_no_stream_goaway_close": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data no stream and no goaway and close session",
												},
												"est_cb_no_tuple": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Est callback no tuple",
												},
												"data_cb_no_tuple": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data callback no tuple",
												},
												"ctx_alloc_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Context alloc fail",
												},
												"stream_alloc_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Stream alloc fail",
												},
												"http_conn_alloc_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP connection allocation fail",
												},
												"request_header_alloc_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Request/Header allocation fail",
												},
												"decompress_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Decompress fail",
												},
												"invalid_frame_size": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid frame size",
												},
												"invalid_version": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid version",
												},
												"compress_ctx_alloc_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Compression context allocation fail",
												},
												"header_compress_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Header compress fail",
												},
												"http_err_stream_closed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP error stream already closed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_http2": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"protocol_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Protocol Error",
												},
												"internal_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Internal Error",
												},
												"proxy_alloc_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP2 Proxy alloc Error",
												},
												"split_buff_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Splitting Buffer Failed",
												},
												"invalid_frame_size": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid Frame Size Rcvd",
												},
												"error_max_invalid_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Max Invalid Stream Rcvd",
												},
												"data_no_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DATA Frame Rcvd on non-existent stream",
												},
												"flow_control_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Flow Control Error",
												},
												"settings_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Settings Timeout",
												},
												"frame_size_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Frame Size Error",
												},
												"refused_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Refused Stream",
												},
												"cancel": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cancel",
												},
												"compression_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for compression error",
												},
												"connect_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for connect error",
												},
												"enhance_your_calm": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for enhance your calm error",
												},
												"inadequate_security": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for inadequate security",
												},
												"http_1_1_required": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP1.1 Required",
												},
												"deflate_alloc_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for deflate alloc fail",
												},
												"inflate_alloc_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for inflate alloc fail",
												},
												"inflate_header_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inflate Header Fail",
												},
												"bad_connection_preface": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Connection Preface",
												},
												"cant_allocate_control_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate control frame",
												},
												"cant_allocate_settings_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate SETTINGS frame",
												},
												"bad_frame_type_for_stream_state": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad frame type for stream state",
												},
												"wrong_stream_state": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Wrong Stream State",
												},
												"data_queue_alloc_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Queue Alloc Error",
												},
												"buff_alloc_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Buff alloc error",
												},
												"cant_allocate_rst_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate RST_STREAM frame",
												},
												"cant_allocate_goaway_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate GOAWAY frame",
												},
												"cant_allocate_ping_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate PING frame",
												},
												"cant_allocate_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate stream",
												},
												"cant_allocate_window_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate WINDOW_UPDATE frame",
												},
												"header_no_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for header no stream",
												},
												"header_padlen_gt_frame_payload": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Header padlen greater than frame payload size",
												},
												"streams_gt_max_concur_streams": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Streams greater than max allowed concurrent streams",
												},
												"idle_state_unexpected_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unxpected frame received in idle state",
												},
												"reserved_local_state_unexpected_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in reserved local state",
												},
												"reserved_remote_state_unexpected_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in reserved remote state",
												},
												"half_closed_remote_state_unexpected_fra": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in half closed remote state",
												},
												"closed_state_unexpected_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in closed state",
												},
												"zero_window_size_on_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Window Update with zero increment rcvd",
												},
												"exceeds_max_window_size_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Window Update with increment that results in exceeding max window",
												},
												"continuation_before_headers": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CONTINUATION frame with no headers frame",
												},
												"invalid_frame_during_headers": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for frame before headers were complete",
												},
												"headers_after_continuation": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for headers frame before CONTINUATION was complete",
												},
												"invalid_push_promise": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unexpected PUSH_PROMISE frame",
												},
												"invalid_stream_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for received invalid stream ID",
												},
												"headers_interleaved": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for headers interleaved on streams",
												},
												"trailers_no_end_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for trailers not marked as end-of-stream",
												},
												"invalid_setting_value": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid setting-frame value",
												},
												"invalid_window_update": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for window-update value out of range",
												},
												"alloc_fail_total": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Alloc Fail - Total",
												},
												"err_rcvd_total": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Rcvd - Total",
												},
												"err_sent_total": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Rent - Total",
												},
												"err_sent_proto_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - PROTOCOL_ERROR",
												},
												"err_sent_internal_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - INTERNAL_ERROR",
												},
												"err_sent_flow_control": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - FLOW_CONTROL_ERROR",
												},
												"err_sent_setting_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - SETTINGS_TIMEOUT",
												},
												"err_sent_stream_closed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - STREAM_CLOSED",
												},
												"err_sent_frame_size_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - FRAME_SIZE_ERROR",
												},
												"err_sent_refused_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - REFUSED_STREAM",
												},
												"err_sent_cancel": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - CANCEL",
												},
												"err_sent_compression_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - COMPRESSION_ERROR",
												},
												"err_sent_connect_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - CONNECT_ERROR",
												},
												"err_sent_your_calm": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - ENHANCE_YOUR_CALM",
												},
												"err_sent_inadequate_security": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - INADEQUATE_SECURITY",
												},
												"err_sent_http11_required": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - HTTP_1_1_REQUIRED",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"protocol_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Protocol Error",
												},
												"internal_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Internal Error",
												},
												"proxy_alloc_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP2 Proxy alloc Error",
												},
												"split_buff_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Splitting Buffer Failed",
												},
												"invalid_frame_size": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid Frame Size Rcvd",
												},
												"error_max_invalid_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Max Invalid Stream Rcvd",
												},
												"data_no_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DATA Frame Rcvd on non-existent stream",
												},
												"flow_control_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Flow Control Error",
												},
												"settings_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Settings Timeout",
												},
												"frame_size_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Frame Size Error",
												},
												"refused_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Refused Stream",
												},
												"cancel": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cancel",
												},
												"compression_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for compression error",
												},
												"connect_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for connect error",
												},
												"enhance_your_calm": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for enhance your calm error",
												},
												"inadequate_security": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for inadequate security",
												},
												"http_1_1_required": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP1.1 Required",
												},
												"deflate_alloc_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for deflate alloc fail",
												},
												"inflate_alloc_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for inflate alloc fail",
												},
												"inflate_header_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inflate Header Fail",
												},
												"bad_connection_preface": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Connection Preface",
												},
												"cant_allocate_control_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate control frame",
												},
												"cant_allocate_settings_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate SETTINGS frame",
												},
												"bad_frame_type_for_stream_state": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad frame type for stream state",
												},
												"wrong_stream_state": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Wrong Stream State",
												},
												"data_queue_alloc_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Queue Alloc Error",
												},
												"buff_alloc_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Buff alloc error",
												},
												"cant_allocate_rst_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate RST_STREAM frame",
												},
												"cant_allocate_goaway_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate GOAWAY frame",
												},
												"cant_allocate_ping_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate PING frame",
												},
												"cant_allocate_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate stream",
												},
												"cant_allocate_window_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate WINDOW_UPDATE frame",
												},
												"header_no_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for header no stream",
												},
												"header_padlen_gt_frame_payload": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Header padlen greater than frame payload size",
												},
												"streams_gt_max_concur_streams": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Streams greater than max allowed concurrent streams",
												},
												"idle_state_unexpected_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unxpected frame received in idle state",
												},
												"reserved_local_state_unexpected_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in reserved local state",
												},
												"reserved_remote_state_unexpected_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in reserved remote state",
												},
												"half_closed_remote_state_unexpected_fra": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in half closed remote state",
												},
												"closed_state_unexpected_frame": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in closed state",
												},
												"zero_window_size_on_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Window Update with zero increment rcvd",
												},
												"exceeds_max_window_size_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Window Update with increment that results in exceeding max window",
												},
												"continuation_before_headers": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CONTINUATION frame with no headers frame",
												},
												"invalid_frame_during_headers": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for frame before headers were complete",
												},
												"headers_after_continuation": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for headers frame before CONTINUATION was complete",
												},
												"invalid_push_promise": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unexpected PUSH_PROMISE frame",
												},
												"invalid_stream_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for received invalid stream ID",
												},
												"headers_interleaved": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for headers interleaved on streams",
												},
												"trailers_no_end_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for trailers not marked as end-of-stream",
												},
												"invalid_setting_value": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid setting-frame value",
												},
												"invalid_window_update": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for window-update value out of range",
												},
												"alloc_fail_total": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Alloc Fail - Total",
												},
												"err_rcvd_total": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Rcvd - Total",
												},
												"err_sent_total": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Rent - Total",
												},
												"err_sent_proto_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - PROTOCOL_ERROR",
												},
												"err_sent_internal_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - INTERNAL_ERROR",
												},
												"err_sent_flow_control": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - FLOW_CONTROL_ERROR",
												},
												"err_sent_setting_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - SETTINGS_TIMEOUT",
												},
												"err_sent_stream_closed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - STREAM_CLOSED",
												},
												"err_sent_frame_size_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - FRAME_SIZE_ERROR",
												},
												"err_sent_refused_stream": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - REFUSED_STREAM",
												},
												"err_sent_cancel": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - CANCEL",
												},
												"err_sent_compression_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - COMPRESSION_ERROR",
												},
												"err_sent_connect_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - CONNECT_ERROR",
												},
												"err_sent_your_calm": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - ENHANCE_YOUR_CALM",
												},
												"err_sent_inadequate_security": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - INADEQUATE_SECURITY",
												},
												"err_sent_http11_required": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - HTTP_1_1_REQUIRED",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_l7session": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"conn_not_exist": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Conn does not exist",
												},
												"wbuf_cb_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Wbuf event callback failed",
												},
												"err_event": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Err event from TCP",
												},
												"err_cb_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Err event callback failed",
												},
												"server_conn_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server connection failed",
												},
												"server_select_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection fail",
												},
												"data_cb_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data event callback fail",
												},
												"hps_fwdreq_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fwd req fail",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"conn_not_exist": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Conn does not exist",
												},
												"wbuf_cb_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Wbuf event callback failed",
												},
												"err_event": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Err event from TCP",
												},
												"err_cb_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Err event callback failed",
												},
												"server_conn_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server connection failed",
												},
												"server_select_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection fail",
												},
												"data_cb_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data event callback fail",
												},
												"hps_fwdreq_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fwd req fail",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_smpp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"msg_proxy_client_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of SMPP messages received from client but failed to forward to server",
												},
												"msg_proxy_server_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of SMPP messages received from server but failed to forward to client",
												},
												"msg_proxy_fail_start_server_conn": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of server connection created failed",
												},
												"select_client_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Select failed",
												},
												"select_server_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fail to select server conn",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"msg_proxy_client_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of SMPP messages received from client but failed to forward to server",
												},
												"msg_proxy_server_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of SMPP messages received from server but failed to forward to client",
												},
												"msg_proxy_fail_start_server_conn": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of server connection created failed",
												},
												"select_client_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Select failed",
												},
												"select_server_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fail to select server conn",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_smtp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"no_proxy": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No proxy error",
												},
												"parse_req_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse request failure",
												},
												"server_select_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
												},
												"forward_req_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Forward request failure",
												},
												"forward_req_data_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Forward REQ data failure",
												},
												"snat_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT failure",
												},
												"send_client_service_not_ready": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Sent client serv-not-rdy",
												},
												"recv_server_unknow_reply_code": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Recv server unknown-code",
												},
												"read_request_line_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Read request line fail",
												},
												"get_all_headers_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Get all headers fail",
												},
												"too_many_headers": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Too many headers",
												},
												"line_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Line too long",
												},
												"line_extend_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Line extend fail",
												},
												"line_table_extend_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Table extend fail",
												},
												"parse_request_line_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse request line fail",
												},
												"insert_resonse_line_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Ins response line fail",
												},
												"remove_resonse_line_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Del response line fail",
												},
												"parse_resonse_line_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse response line fail",
												},
												"server_starttls_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server side STARTTLS fail",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"no_proxy": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No proxy error",
												},
												"parse_req_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse request failure",
												},
												"server_select_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
												},
												"forward_req_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Forward request failure",
												},
												"forward_req_data_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Forward REQ data failure",
												},
												"snat_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT failure",
												},
												"send_client_service_not_ready": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Sent client serv-not-rdy",
												},
												"recv_server_unknow_reply_code": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Recv server unknown-code",
												},
												"read_request_line_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Read request line fail",
												},
												"get_all_headers_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Get all headers fail",
												},
												"too_many_headers": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Too many headers",
												},
												"line_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Line too long",
												},
												"line_extend_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Line extend fail",
												},
												"line_table_extend_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Table extend fail",
												},
												"parse_request_line_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse request line fail",
												},
												"insert_resonse_line_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Ins response line fail",
												},
												"remove_resonse_line_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Del response line fail",
												},
												"parse_resonse_line_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse response line fail",
												},
												"server_starttls_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server side STARTTLS fail",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_mqtt": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parse_connect_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse connect failure",
												},
												"parse_publish_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse publish failure",
												},
												"parse_subscribe_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse subscribe failure",
												},
												"parse_unsubscribe_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse unsubscribe failure",
												},
												"tuple_not_linked": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tuple-not-linked failure",
												},
												"tuple_already_linked": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tuple-already-linked failure",
												},
												"conn_null": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Null conn",
												},
												"client_id_null": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Null client id",
												},
												"session_exist": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Session already exist",
												},
												"insertion_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Insertion failure",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"parse_connect_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse connect failure",
												},
												"parse_publish_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse publish failure",
												},
												"parse_subscribe_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse subscribe failure",
												},
												"parse_unsubscribe_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse unsubscribe failure",
												},
												"tuple_not_linked": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tuple-not-linked failure",
												},
												"tuple_already_linked": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for tuple-already-linked failure",
												},
												"conn_null": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Null conn",
												},
												"client_id_null": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Null client id",
												},
												"session_exist": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Session already exist",
												},
												"insertion_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Insertion failure",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_icap": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"app_serv_conn_no_pcb_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for App Server Conn no ES PCB Err Stats",
												},
												"app_serv_conn_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for App Server Conn Err Stats",
												},
												"chunk1_hdr_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Chunk Hdr Err1 Stats",
												},
												"chunk2_hdr_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Chunk Hdr Err2 Stats",
												},
												"chunk_bad_trail_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Chunk Bad Trail Err Stats",
												},
												"no_payload_next_buff_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Payload In Next Buff Err Stats",
												},
												"no_payload_buff_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Payload Buff Err Stats",
												},
												"resp_hdr_incomplete_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Hdr Incomplete Err Stats",
												},
												"serv_sel_fail_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server Select Fail Err Stats",
												},
												"start_icap_conn_fail_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Start ICAP conn fail Stats",
												},
												"prep_req_fail_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Prepare ICAP req fail Err Stats",
												},
												"icap_ver_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICAP Ver Err Stats",
												},
												"icap_line_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICAP Line Err Stats",
												},
												"encap_hdr_incomplete_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Encap HDR Incomplete Err Stats",
												},
												"no_icap_resp_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No ICAP Resp Err Stats",
												},
												"resp_line_read_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Line Read Err Stats",
												},
												"resp_line_parse_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Line Parse Err Stats",
												},
												"resp_hdr_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Hdr Err Stats",
												},
												"req_hdr_incomplete_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Req Hdr Incomplete Err Stats",
												},
												"no_status_code_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Status Code Err Stats",
												},
												"http_resp_line_read_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP Response Line Read Err Stats",
												},
												"http_resp_line_parse_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP Response Line Parse Err Stats",
												},
												"http_resp_hdr_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP Resp Hdr Err Stats",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"app_serv_conn_no_pcb_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for App Server Conn no ES PCB Err Stats",
												},
												"app_serv_conn_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for App Server Conn Err Stats",
												},
												"chunk1_hdr_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Chunk Hdr Err1 Stats",
												},
												"chunk2_hdr_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Chunk Hdr Err2 Stats",
												},
												"chunk_bad_trail_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Chunk Bad Trail Err Stats",
												},
												"no_payload_next_buff_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Payload In Next Buff Err Stats",
												},
												"no_payload_buff_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Payload Buff Err Stats",
												},
												"resp_hdr_incomplete_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Hdr Incomplete Err Stats",
												},
												"serv_sel_fail_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server Select Fail Err Stats",
												},
												"start_icap_conn_fail_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Start ICAP conn fail Stats",
												},
												"prep_req_fail_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Prepare ICAP req fail Err Stats",
												},
												"icap_ver_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICAP Ver Err Stats",
												},
												"icap_line_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICAP Line Err Stats",
												},
												"encap_hdr_incomplete_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Encap HDR Incomplete Err Stats",
												},
												"no_icap_resp_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No ICAP Resp Err Stats",
												},
												"resp_line_read_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Line Read Err Stats",
												},
												"resp_line_parse_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Line Parse Err Stats",
												},
												"resp_hdr_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Hdr Err Stats",
												},
												"req_hdr_incomplete_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Req Hdr Incomplete Err Stats",
												},
												"no_status_code_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Status Code Err Stats",
												},
												"http_resp_line_read_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP Response Line Read Err Stats",
												},
												"http_resp_line_parse_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP Response Line Parse Err Stats",
												},
												"http_resp_hdr_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP Resp Hdr Err Stats",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_sip": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"msg_proxy_client_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of SIP messages received from client but failed to forward to server",
												},
												"msg_proxy_server_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of SIP messages received from server but failed to forward to client",
												},
												"msg_proxy_fail_start_server_conn": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of server connection create failed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"msg_proxy_client_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of SIP messages received from client but failed to forward to server",
												},
												"msg_proxy_server_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of SIP messages received from server but failed to forward to client",
												},
												"msg_proxy_fail_start_server_conn": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of server connection create failed",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_hw_compress": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"failure_count": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total failure count",
												},
												"failure_code": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Last failure code",
												},
												"ring_full_count": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Compression queue full",
												},
												"max_outstanding_request_count": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Max queued request count",
												},
												"max_outstanding_submit_count": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Max queued submit count",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"failure_count": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total failure count",
												},
												"failure_code": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Last failure code",
												},
												"ring_full_count": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Compression queue full",
												},
												"max_outstanding_request_count": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Max queued request count",
												},
												"max_outstanding_submit_count": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Max queued submit count",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_mysql": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"session_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Session err",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"session_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Session err",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_mssql": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"session_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Session err",
												},
												"auth_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Authentication Failure",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"session_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Session err",
												},
												"auth_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Authentication Failure",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_crl_srcip": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"out_of_sessions": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of sessions",
												},
												"too_many_sessions": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Too many sessions consumed",
												},
												"threshold_exceed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Threshold exceeded count",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"out_of_sessions": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of sessions",
												},
												"too_many_sessions": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Too many sessions consumed",
												},
												"threshold_exceed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Threshold exceeded count",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_generic": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"svrsel_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of server selection failed",
												},
												"no_route": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of no routes",
												},
												"snat_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of snat failures",
												},
												"client_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of client failures",
												},
												"server_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of server failures",
												},
												"mismatch_fwd_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Diameter mismatch fwd session id",
												},
												"mismatch_rev_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Diameter mismatch rev session id",
												},
												"unkwn_cmd_code": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Diameter unkown cmd code",
												},
												"no_session_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Diameter no session id avp",
												},
												"no_fwd_tuple": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Diameter no fwd tuple matched",
												},
												"no_rev_tuple": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Diameter no rev tuple matched",
												},
												"dcmsg_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Diameter cross cpu error",
												},
												"retry_client_request_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Diameter retry client request fail",
												},
												"reply_unknown_session_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Reply with unknown session ID error info",
												},
												"client_select_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fail to select client",
												},
												"invalid_avp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for AVP value contains illegal chars",
												},
												"reply_error_info_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fail to reply error info to peer",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"svrsel_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of server selection failed",
												},
												"no_route": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of no routes",
												},
												"snat_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of snat failures",
												},
												"client_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of client failures",
												},
												"server_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Number of server failures",
												},
												"mismatch_fwd_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Diameter mismatch fwd session id",
												},
												"mismatch_rev_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Diameter mismatch rev session id",
												},
												"unkwn_cmd_code": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Diameter unkown cmd code",
												},
												"no_session_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Diameter no session id avp",
												},
												"no_fwd_tuple": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Diameter no fwd tuple matched",
												},
												"no_rev_tuple": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Diameter no rev tuple matched",
												},
												"dcmsg_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Diameter cross cpu error",
												},
												"retry_client_request_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Diameter retry client request fail",
												},
												"reply_unknown_session_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Reply with unknown session ID error info",
												},
												"client_select_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fail to select client",
												},
												"invalid_avp": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for AVP value contains illegal chars",
												},
												"reply_error_info_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fail to reply error info to peer",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_persist": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hash_tbl_trylock_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl lock fail",
												},
												"hash_tbl_create_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl create fail",
												},
												"hash_tbl_rst_updown": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl reset (up/down)",
												},
												"hash_tbl_rst_adddel": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl reset (add/del)",
												},
												"url_hash_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for URL hash persist fail",
												},
												"header_hash_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Header hash persist fail",
												},
												"src_ip_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP persist fail",
												},
												"src_ip_new_sess_cache_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP new sess fail (c)",
												},
												"src_ip_new_sess_sel_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP new sess fail (s)",
												},
												"src_ip_hash_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP hash persist fail",
												},
												"dst_ip_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP persist fail",
												},
												"dst_ip_new_sess_cache_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP new sess fail (c)",
												},
												"dst_ip_new_sess_sel_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP new sess fail (s)",
												},
												"dst_ip_hash_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP hash persist fail",
												},
												"cssl_sid_not_found": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client SSL SID not found",
												},
												"cssl_sid_not_match": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client SSL SID not match",
												},
												"sssl_sid_not_found": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server SSL SID not found",
												},
												"sssl_sid_not_match": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server SSL SID not match",
												},
												"ssl_sid_persist_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SSL SID persist fail",
												},
												"ssl_sid_session_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Create SSL SID fail",
												},
												"cookie_persist_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cookie persist fail",
												},
												"cookie_not_found": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Persist cookie not found",
												},
												"cookie_invalid": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid persist cookie",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"hash_tbl_trylock_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl lock fail",
												},
												"hash_tbl_create_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl create fail",
												},
												"hash_tbl_rst_updown": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl reset (up/down)",
												},
												"hash_tbl_rst_adddel": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl reset (add/del)",
												},
												"url_hash_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for URL hash persist fail",
												},
												"header_hash_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Header hash persist fail",
												},
												"src_ip_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP persist fail",
												},
												"src_ip_new_sess_cache_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP new sess fail (c)",
												},
												"src_ip_new_sess_sel_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP new sess fail (s)",
												},
												"src_ip_hash_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP hash persist fail",
												},
												"dst_ip_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP persist fail",
												},
												"dst_ip_new_sess_cache_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP new sess fail (c)",
												},
												"dst_ip_new_sess_sel_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP new sess fail (s)",
												},
												"dst_ip_hash_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP hash persist fail",
												},
												"cssl_sid_not_found": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client SSL SID not found",
												},
												"cssl_sid_not_match": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client SSL SID not match",
												},
												"sssl_sid_not_found": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server SSL SID not found",
												},
												"sssl_sid_not_match": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server SSL SID not match",
												},
												"ssl_sid_persist_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SSL SID persist fail",
												},
												"ssl_sid_session_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Create SSL SID fail",
												},
												"cookie_persist_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cookie persist fail",
												},
												"cookie_not_found": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Persist cookie not found",
												},
												"cookie_invalid": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid persist cookie",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_http_proxy": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parsereq_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse req fail",
												},
												"svrsel_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection fail",
												},
												"fwdreq_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fwd req fail",
												},
												"fwdreqdata_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for fwdreqdata_fail",
												},
												"snat_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT failure",
												},
												"req_over_limit": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Request over limit",
												},
												"req_rate_over_limit": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Request rate over limit",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"parsereq_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse req fail",
												},
												"svrsel_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection fail",
												},
												"fwdreq_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fwd req fail",
												},
												"fwdreqdata_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for fwdreqdata_fail",
												},
												"snat_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT failure",
												},
												"req_over_limit": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Request over limit",
												},
												"req_rate_over_limit": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Request rate over limit",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_l4": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"syncookiessentfailed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP SYN cookie snt fail",
												},
												"svrselfail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server sel failure",
												},
												"snat_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT failure",
												},
												"snat_no_fwd_route": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT no fwd route",
												},
												"snat_no_rev_route": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT no rev route",
												},
												"snat_icmp_error_process": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT ICMP Process",
												},
												"snat_icmp_no_match": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT ICMP No Match",
												},
												"smart_nat_id_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Auto NAT id mismatch",
												},
												"syncookiescheckfailed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP SYN cookie failed",
												},
												"connlimit_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Conn Limit drops",
												},
												"conn_rate_limit_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Conn rate limit drops",
												},
												"conn_rate_limit_reset": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Conn rate limit resets",
												},
												"dns_policy_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DNS Policy Drop",
												},
												"no_resourse_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No resource drop",
												},
												"bw_rate_limit_exceed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for BW-Limit Exceed drop",
												},
												"l4_cps_exceed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 CPS exceed drop",
												},
												"nat_cps_exceed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT CPS exceed drop",
												},
												"l7_cps_exceed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L7 CPS exceed drop",
												},
												"ssl_cps_exceed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SSL CPS exceed drop",
												},
												"ssl_tpt_exceed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SSL TPT exceed drop",
												},
												"concurrent_conn_exceed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3V Conn Limit Drop",
												},
												"svr_syn_handshake_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 server handshake fail",
												},
												"synattack": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 SYN attack",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"syncookiessentfailed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP SYN cookie snt fail",
												},
												"svrselfail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server sel failure",
												},
												"snat_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT failure",
												},
												"snat_no_fwd_route": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT no fwd route",
												},
												"snat_no_rev_route": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT no rev route",
												},
												"snat_icmp_error_process": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT ICMP Process",
												},
												"snat_icmp_no_match": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT ICMP No Match",
												},
												"smart_nat_id_mismatch": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Auto NAT id mismatch",
												},
												"syncookiescheckfailed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP SYN cookie failed",
												},
												"connlimit_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Conn Limit drops",
												},
												"conn_rate_limit_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Conn rate limit drops",
												},
												"conn_rate_limit_reset": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Conn rate limit resets",
												},
												"dns_policy_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DNS Policy Drop",
												},
												"no_resourse_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No resource drop",
												},
												"bw_rate_limit_exceed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for BW-Limit Exceed drop",
												},
												"l4_cps_exceed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 CPS exceed drop",
												},
												"nat_cps_exceed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NAT CPS exceed drop",
												},
												"l7_cps_exceed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L7 CPS exceed drop",
												},
												"ssl_cps_exceed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SSL CPS exceed drop",
												},
												"ssl_tpt_exceed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SSL TPT exceed drop",
												},
												"concurrent_conn_exceed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3V Conn Limit Drop",
												},
												"svr_syn_handshake_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 server handshake fail",
												},
												"synattack": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L4 SYN attack",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_fast_http": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parsereq_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse req fail",
												},
												"svrsel_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection fail",
												},
												"fwdreq_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fwd req fail",
												},
												"fwdreqdata_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fwd req data fail",
												},
												"snat_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT failure",
												},
												"full_proxy_fpga_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Full proxy fpga err",
												},
												"req_over_limit": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Request over limit",
												},
												"req_rate_over_limit": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Request rate over limit",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"parsereq_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse req fail",
												},
												"svrsel_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection fail",
												},
												"fwdreq_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fwd req fail",
												},
												"fwdreqdata_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Fwd req data fail",
												},
												"snat_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT failure",
												},
												"full_proxy_fpga_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Full proxy fpga err",
												},
												"req_over_limit": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Request over limit",
												},
												"req_rate_over_limit": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Request rate over limit",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_ftp_proxy": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"svrsel_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
												},
												"no_route": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no route failure",
												},
												"snat_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for source nat failure",
												},
												"line_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for line too long",
												},
												"invalid_start_line": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid start line",
												},
												"smp_create_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for smp create fail",
												},
												"data_server_conn_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data svr conn fail",
												},
												"data_send_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data send fail",
												},
												"unsupported_pbsz_value": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PBSZ",
												},
												"unsupported_prot_value": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PROT",
												},
												"unsupported_command": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported cmd",
												},
												"bad_sequence": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Sequence",
												},
												"rsv_persist_conn_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel Persist fail",
												},
												"smp_v6_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv6 fail",
												},
												"smp_v4_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv4 fail",
												},
												"insert_tuple_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel insert tuple fail",
												},
												"cl_est_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client EST state erro",
												},
												"ser_connecting_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv CTNG state error",
												},
												"server_response_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv RESP state error",
												},
												"cl_request_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client RQ state error",
												},
												"data_conn_start_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Start state error",
												},
												"data_serv_connecting_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Serv CTNG error",
												},
												"data_serv_connected_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Serv CTED error",
												},
												"auth_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Auth Failure",
												},
												"ds_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Host Domain Name isn't resolved",
												},
												"cant_find_port": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cant find port",
												},
												"cant_find_eprt": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cant find eprt",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"svrsel_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
												},
												"no_route": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no route failure",
												},
												"snat_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for source nat failure",
												},
												"line_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for line too long",
												},
												"invalid_start_line": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid start line",
												},
												"smp_create_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for smp create fail",
												},
												"data_server_conn_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data svr conn fail",
												},
												"data_send_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data send fail",
												},
												"unsupported_pbsz_value": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PBSZ",
												},
												"unsupported_prot_value": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PROT",
												},
												"unsupported_command": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported cmd",
												},
												"bad_sequence": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Sequence",
												},
												"rsv_persist_conn_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel Persist fail",
												},
												"smp_v6_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv6 fail",
												},
												"smp_v4_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv4 fail",
												},
												"insert_tuple_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel insert tuple fail",
												},
												"cl_est_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client EST state erro",
												},
												"ser_connecting_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv CTNG state error",
												},
												"server_response_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv RESP state error",
												},
												"cl_request_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client RQ state error",
												},
												"data_conn_start_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Start state error",
												},
												"data_serv_connecting_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Serv CTNG error",
												},
												"data_serv_connected_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Serv CTED error",
												},
												"auth_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Auth Failure",
												},
												"ds_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Host Domain Name isn't resolved",
												},
												"cant_find_port": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cant find port",
												},
												"cant_find_eprt": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cant find eprt",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_imap_proxy": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"svrsel_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
												},
												"no_route": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no route failure",
												},
												"snat_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for source nat failure",
												},
												"line_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for line too long",
												},
												"invalid_start_line": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid start line",
												},
												"cant_find_pasv": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cant find pasv",
												},
												"smp_create_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for smp create fail",
												},
												"data_server_conn_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data svr conn fail",
												},
												"data_send_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data send fail",
												},
												"cant_find_epsv": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cant find epsv",
												},
												"auth_unsupported": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported auth",
												},
												"unsupported_pbsz_value": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PBSZ",
												},
												"unsupported_prot_value": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PROT",
												},
												"bad_sequence": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Sequence",
												},
												"rsv_persist_conn_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel Persist fail",
												},
												"smp_v6_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv6 fail",
												},
												"smp_v4_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv4 fail",
												},
												"insert_tuple_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel insert tuple fail",
												},
												"cl_est_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client EST state erro",
												},
												"ser_connecting_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv CTNG state error",
												},
												"server_response_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv RESP state error",
												},
												"cl_request_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client RQ state error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"svrsel_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
												},
												"no_route": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no route failure",
												},
												"snat_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for source nat failure",
												},
												"line_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for line too long",
												},
												"invalid_start_line": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid start line",
												},
												"cant_find_pasv": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cant find pasv",
												},
												"smp_create_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for smp create fail",
												},
												"data_server_conn_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data svr conn fail",
												},
												"data_send_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data send fail",
												},
												"cant_find_epsv": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cant find epsv",
												},
												"auth_unsupported": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported auth",
												},
												"unsupported_pbsz_value": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PBSZ",
												},
												"unsupported_prot_value": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PROT",
												},
												"bad_sequence": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Sequence",
												},
												"rsv_persist_conn_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel Persist fail",
												},
												"smp_v6_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv6 fail",
												},
												"smp_v4_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv4 fail",
												},
												"insert_tuple_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel insert tuple fail",
												},
												"cl_est_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client EST state erro",
												},
												"ser_connecting_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv CTNG state error",
												},
												"server_response_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv RESP state error",
												},
												"cl_request_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client RQ state error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_pop3_proxy": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"svrsel_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
												},
												"no_route": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no route failure",
												},
												"snat_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for source nat failure",
												},
												"line_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for line too long",
												},
												"invalid_start_line": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid start line",
												},
												"unsupported_command": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported cmd",
												},
												"bad_sequence": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Sequence",
												},
												"rsv_persist_conn_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel Persist fail",
												},
												"smp_v6_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv6 fail",
												},
												"smp_v4_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv4 fail",
												},
												"insert_tuple_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel insert tuple fail",
												},
												"cl_est_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client EST state erro",
												},
												"ser_connecting_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv CTNG state error",
												},
												"server_response_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv RESP state error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"svrsel_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
												},
												"no_route": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no route failure",
												},
												"snat_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for source nat failure",
												},
												"line_too_long": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for line too long",
												},
												"invalid_start_line": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid start line",
												},
												"unsupported_command": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported cmd",
												},
												"bad_sequence": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Sequence",
												},
												"rsv_persist_conn_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel Persist fail",
												},
												"smp_v6_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv6 fail",
												},
												"smp_v4_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv4 fail",
												},
												"insert_tuple_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel insert tuple fail",
												},
												"cl_est_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client EST state erro",
												},
												"ser_connecting_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv CTNG state error",
												},
												"server_response_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv RESP state error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_switch": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lacp_tx_intf_err_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for LACP interface error corrected",
												},
												"unnumbered_nat_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unnumbered NAT error",
												},
												"unnumbered_unsupported_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported protocol for unnumbered",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"lacp_tx_intf_err_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for LACP interface error corrected",
												},
												"unnumbered_nat_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unnumbered NAT error",
												},
												"unnumbered_unsupported_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported protocol for unnumbered",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_rc_cache": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"rv_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Revalidation Failures",
												},
												"content_toobig": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Policy Content Too Big",
												},
												"content_toosmall": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Policy Content Too Small",
												},
												"entry_create_failures": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Entry Create failures",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"rv_failure": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Revalidation Failures",
												},
												"content_toobig": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Policy Content Too Big",
												},
												"content_toosmall": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Policy Content Too Small",
												},
												"entry_create_failures": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Entry Create failures",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"so_counters": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"so_pkts_slb_nat_reserve_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total SLB NAT reserve failures",
												},
												"so_pkts_slb_nat_release_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total SLB NAT release failures",
												},
												"so_pkts_l2redirect_dest_mac_zero_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Destination MAC Address zero Drop",
												},
												"so_pkts_l2redirect_interface_not_up": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2redirect Intf is not UP",
												},
												"so_pkts_l2redirect_invalid_redirect_inf": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Redirect Table Error due to invalid redirect info",
												},
												"so_pkts_l3_redirect_encap_error_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect encap error drop during transmission",
												},
												"so_pkts_l3_redirect_inner_mac_zero_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect inner mac zero drop during transmission",
												},
												"so_pkts_l3_redirect_table_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect Table error Drop",
												},
												"so_pkts_l3_redirect_fragmentation_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 redirect encap Fragmentation error",
												},
												"so_pkts_l3_redirect_table_no_entry_foun": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 redirect Table no redirect entry found error",
												},
												"so_pkts_l3_redirect_invalid_dev_dir": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect Invalid Device direction during transmission",
												},
												"so_pkts_l3_redirect_chassis_dest_mac_er": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect RX multi-slot Destination MAC Error",
												},
												"so_pkts_l2redirect_vlan_retrieval_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2 redirect pkt vlan not retrieved",
												},
												"so_pkts_l2redirect_port_retrieval_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2 redirect pkt port not retrieved",
												},
												"so_pkts_l2redirect_loop_detect_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2 redirect pkt loop detected and dropped",
												},
												"so_pkts_l2redirect_frag_vlan_retrieval_": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2 redirect pkt frag vlan not retrieved",
												},
												"so_pkts_l2redirect_tx_frag_vlan_add_fai": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2 redirect tx pkt failed to add vlan",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"so_pkts_slb_nat_reserve_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total SLB NAT reserve failures",
												},
												"so_pkts_slb_nat_release_fail": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total SLB NAT release failures",
												},
												"so_pkts_l2redirect_dest_mac_zero_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Destination MAC Address zero Drop",
												},
												"so_pkts_l2redirect_interface_not_up": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2redirect Intf is not UP",
												},
												"so_pkts_l2redirect_invalid_redirect_inf": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Redirect Table Error due to invalid redirect info",
												},
												"so_pkts_l3_redirect_encap_error_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect encap error drop during transmission",
												},
												"so_pkts_l3_redirect_inner_mac_zero_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect inner mac zero drop during transmission",
												},
												"so_pkts_l3_redirect_table_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect Table error Drop",
												},
												"so_pkts_l3_redirect_fragmentation_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 redirect encap Fragmentation error",
												},
												"so_pkts_l3_redirect_table_no_entry_foun": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 redirect Table no redirect entry found error",
												},
												"so_pkts_l3_redirect_invalid_dev_dir": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect Invalid Device direction during transmission",
												},
												"so_pkts_l3_redirect_chassis_dest_mac_er": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L3 Redirect RX multi-slot Destination MAC Error",
												},
												"so_pkts_l2redirect_vlan_retrieval_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2 redirect pkt vlan not retrieved",
												},
												"so_pkts_l2redirect_port_retrieval_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2 redirect pkt port not retrieved",
												},
												"so_pkts_l2redirect_loop_detect_drop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2 redirect pkt loop detected and dropped",
												},
												"so_pkts_l2redirect_frag_vlan_retrieval_": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2 redirect pkt frag vlan not retrieved",
												},
												"so_pkts_l2redirect_tx_frag_vlan_add_fai": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for L2 redirect tx pkt failed to add vlan",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_plyr_id_gbl": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"total_invalid_playerid_pkts": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid playerid packets",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"total_invalid_playerid_pkts": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid playerid packets",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_sport_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"total_reset": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total rate exceed reset",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"total_reset": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total rate exceed reset",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"logging_local_log_global": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enqueue_full": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total local-log queue full",
												},
												"enqueue_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total local-log enqueue error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"enqueue_full": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total local-log queue full",
												},
												"enqueue_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total local-log enqueue error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_mlb": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mlb_dcmsg_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Dcmsg error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"mlb_dcmsg_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Dcmsg error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_link_probe": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"err_entry_create_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Entry Creation Failure",
												},
												"err_entry_create_oom": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Entry Creation Out of Memory",
												},
												"err_entry_insert_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Entry Insert Failed",
												},
												"err_tmpl_probe_create_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Probe Template Creation Failure",
												},
												"err_tmpl_probe_create_oom": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Probe Template Creation Out of Memory",
												},
												"total_http_response_bad": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total HTTP responses not matching probe template config",
												},
												"total_tcp_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP errors in probes sent out",
												},
												"err_smart_nat_alloc": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error creating Smart NAT Instance",
												},
												"err_smart_nat_port_alloc": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error obtaining Smart NAT source port",
												},
												"err_l4_sess_alloc": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error allocating L4 session for probe",
												},
												"err_probe_tcp_conn_send": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error in initiating TCP connection for probe",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"err_entry_create_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Entry Creation Failure",
												},
												"err_entry_create_oom": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Entry Creation Out of Memory",
												},
												"err_entry_insert_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Entry Insert Failed",
												},
												"err_tmpl_probe_create_failed": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Probe Template Creation Failure",
												},
												"err_tmpl_probe_create_oom": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Probe Template Creation Out of Memory",
												},
												"total_http_response_bad": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total HTTP responses not matching probe template config",
												},
												"total_tcp_err": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total TCP errors in probes sent out",
												},
												"err_smart_nat_alloc": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error creating Smart NAT Instance",
												},
												"err_smart_nat_port_alloc": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error obtaining Smart NAT source port",
												},
												"err_l4_sess_alloc": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error allocating L4 session for probe",
												},
												"err_probe_tcp_conn_send": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error in initiating TCP connection for probe",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"slb_rpz": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"trigger_stats_inc": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"set_bw_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RPZ Set Class-list Error",
												},
												"parse_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RPZ Parse Error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"trigger_stats_rate": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold_exceeded_by": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
												},
												"duration": {
													Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
												},
												"set_bw_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RPZ Set Class-list Error",
												},
												"parse_error": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total RPZ Parse Error",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
			"trigger_sys_obj_stats_severity": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
						},
						"error_alert": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
						},
						"error_warning": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
						},
						"error_critical": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
						},
						"drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
						},
						"drop_alert": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
						},
						"drop_warning": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
						},
						"drop_critical": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChange2479(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChange2479 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChange2479
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SystemCtrLibAcct = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemCtrLibAcct2480(in["system_ctr_lib_acct"].([]interface{}))
		ret.SystemHardwareAccelerate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerate2483(in["system_hardware_accelerate"].([]interface{}))
		ret.SystemRadiusServer = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemRadiusServer2486(in["system_radius_server"].([]interface{}))
		ret.SystemIpThreatList = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemIpThreatList2489(in["system_ip_threat_list"].([]interface{}))
		ret.SystemFpgaDrop = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop2492(in["system_fpga_drop"].([]interface{}))
		ret.SystemDpdkStats = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats2495(in["system_dpdk_stats"].([]interface{}))
		ret.IpAnomalyDrop = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDrop2498(in["ip_anomaly_drop"].([]interface{}))
		ret.AamAuthenticationGlobal = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthenticationGlobal2501(in["aam_authentication_global"].([]interface{}))
		ret.AamRdns = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamRdns2504(in["aam_rdns"].([]interface{}))
		ret.AamAuthServerLdap = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerLdap2507(in["aam_auth_server_ldap"].([]interface{}))
		ret.AamAuthServerOcsp = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerOcsp2510(in["aam_auth_server_ocsp"].([]interface{}))
		ret.AamAuthServerRadius = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerRadius2513(in["aam_auth_server_radius"].([]interface{}))
		ret.AamAuthServerWin = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin2516(in["aam_auth_server_win"].([]interface{}))
		ret.AamAuthAccount = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthAccount2519(in["aam_auth_account"].([]interface{}))
		ret.AamAuthSamlGlobal = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthSamlGlobal2522(in["aam_auth_saml_global"].([]interface{}))
		ret.AamAuthRelayKerberos = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthRelayKerberos2525(in["aam_auth_relay_kerberos"].([]interface{}))
		ret.AamAuthCaptcha = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthCaptcha2528(in["aam_auth_captcha"].([]interface{}))
		ret.SlbSslError = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslError2531(in["slb_ssl_error"].([]interface{}))
		ret.SlbSslCertRevoke = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke2534(in["slb_ssl_cert_revoke"].([]interface{}))
		ret.SlbSslForwardProxy = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxy2537(in["slb_ssl_forward_proxy"].([]interface{}))
		ret.VpnError = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError2540(in["vpn_error"].([]interface{}))
		ret.Cgnv6DdosProc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProc2543(in["cgnv6_ddos_proc"].([]interface{}))
		ret.Cgnv6Lsn = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Lsn2546(in["cgnv6_lsn"].([]interface{}))
		ret.Cgnv6LsnAlgEsp = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgEsp2549(in["cgnv6_lsn_alg_esp"].([]interface{}))
		ret.Cgnv6LsnAlgPptp = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgPptp2552(in["cgnv6_lsn_alg_pptp"].([]interface{}))
		ret.Cgnv6LsnAlgRtsp = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgRtsp2555(in["cgnv6_lsn_alg_rtsp"].([]interface{}))
		ret.Cgnv6LsnAlgSip = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSip2558(in["cgnv6_lsn_alg_sip"].([]interface{}))
		ret.Cgnv6LsnAlgMgcp = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgMgcp2561(in["cgnv6_lsn_alg_mgcp"].([]interface{}))
		ret.Cgnv6LsnAlgH323 = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgH3232564(in["cgnv6_lsn_alg_h323"].([]interface{}))
		ret.Cgnv6LsnRadius = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRadius2567(in["cgnv6_lsn_radius"].([]interface{}))
		ret.Cgnv6Nat64Global = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64Global2570(in["cgnv6_nat64_global"].([]interface{}))
		ret.Cgnv6DsLiteGlobal = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DsLiteGlobal2573(in["cgnv6_ds_lite_global"].([]interface{}))
		ret.Cgnv6FixedNatGlobal = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobal2576(in["cgnv6_fixed_nat_global"].([]interface{}))
		ret.Cgnv6FixedNatAlgPptp = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgPptp2579(in["cgnv6_fixed_nat_alg_pptp"].([]interface{}))
		ret.Cgnv6FixedNatAlgRtsp = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgRtsp2582(in["cgnv6_fixed_nat_alg_rtsp"].([]interface{}))
		ret.Cgnv6FixedNatAlgSip = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgSip2585(in["cgnv6_fixed_nat_alg_sip"].([]interface{}))
		ret.Cgnv6Pcp = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Pcp2588(in["cgnv6_pcp"].([]interface{}))
		ret.Cgnv6Logging = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Logging2591(in["cgnv6_logging"].([]interface{}))
		ret.Cgnv6L4 = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6L42594(in["cgnv6_l4"].([]interface{}))
		ret.Cgnv6Icmp = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Icmp2597(in["cgnv6_icmp"].([]interface{}))
		ret.Cgnv6HttpAlg = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlg2600(in["cgnv6_http_alg"].([]interface{}))
		ret.Cgnv6Dns64 = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dns642603(in["cgnv6_dns64"].([]interface{}))
		ret.Cgnv6Dhcpv6 = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dhcpv62606(in["cgnv6_dhcpv6"].([]interface{}))
		ret.FwLogging = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwLogging2609(in["fw_logging"].([]interface{}))
		ret.FwGlobal = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGlobal2612(in["fw_global"].([]interface{}))
		ret.FwAlgRtsp = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgRtsp2615(in["fw_alg_rtsp"].([]interface{}))
		ret.FwAlgPptp = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptp2618(in["fw_alg_pptp"].([]interface{}))
		ret.FwRadServer = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServer2621(in["fw_rad_server"].([]interface{}))
		ret.FwTcpSynCookie = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookie2624(in["fw_tcp_syn_cookie"].([]interface{}))
		ret.FwDdosProtection = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwDdosProtection2627(in["fw_ddos_protection"].([]interface{}))
		ret.FwGtp = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp2630(in["fw_gtp"].([]interface{}))
		ret.SystemTcp = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemTcp2633(in["system_tcp"].([]interface{}))
		ret.SlbConnReuse = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbConnReuse2636(in["slb_conn_reuse"].([]interface{}))
		ret.SlbAflow = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbAflow2639(in["slb_aflow"].([]interface{}))
		ret.SlbFix = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFix2642(in["slb_fix"].([]interface{}))
		ret.SlbSpdyProxy = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxy2645(in["slb_spdy_proxy"].([]interface{}))
		ret.SlbHttp2 = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp22648(in["slb_http2"].([]interface{}))
		ret.SlbL7session = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7session2651(in["slb_l7session"].([]interface{}))
		ret.SlbSmpp = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmpp2654(in["slb_smpp"].([]interface{}))
		ret.SlbSmtp = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmtp2657(in["slb_smtp"].([]interface{}))
		ret.SlbMqtt = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMqtt2660(in["slb_mqtt"].([]interface{}))
		ret.SlbIcap = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcap2663(in["slb_icap"].([]interface{}))
		ret.SlbSip = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSip2666(in["slb_sip"].([]interface{}))
		ret.SlbHwCompress = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHwCompress2669(in["slb_hw_compress"].([]interface{}))
		ret.SlbMysql = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMysql2672(in["slb_mysql"].([]interface{}))
		ret.SlbMssql = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMssql2675(in["slb_mssql"].([]interface{}))
		ret.SlbCrlSrcip = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbCrlSrcip2678(in["slb_crl_srcip"].([]interface{}))
		ret.SlbGeneric = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGeneric2681(in["slb_generic"].([]interface{}))
		ret.SlbPersist = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist2684(in["slb_persist"].([]interface{}))
		ret.SlbHttpProxy = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttpProxy2687(in["slb_http_proxy"].([]interface{}))
		ret.SlbL4 = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL42690(in["slb_l4"].([]interface{}))
		ret.SlbFastHttp = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFastHttp2693(in["slb_fast_http"].([]interface{}))
		ret.SlbFtpProxy = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFtpProxy2696(in["slb_ftp_proxy"].([]interface{}))
		ret.SlbImapProxy = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy2699(in["slb_imap_proxy"].([]interface{}))
		ret.SlbPop3Proxy = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPop3Proxy2702(in["slb_pop3_proxy"].([]interface{}))
		ret.SlbSwitch = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSwitch2705(in["slb_switch"].([]interface{}))
		ret.SlbRcCache = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRcCache2708(in["slb_rc_cache"].([]interface{}))
		ret.SoCounters = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCounters2711(in["so_counters"].([]interface{}))
		ret.SlbPlyrIdGbl = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPlyrIdGbl2714(in["slb_plyr_id_gbl"].([]interface{}))
		ret.SlbSportRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSportRate2717(in["slb_sport_rate"].([]interface{}))
		ret.LoggingLocalLogGlobal = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeLoggingLocalLogGlobal2720(in["logging_local_log_global"].([]interface{}))
		ret.SlbMlb = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlb2723(in["slb_mlb"].([]interface{}))
		ret.SlbLinkProbe = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbLinkProbe2726(in["slb_link_probe"].([]interface{}))
		ret.SlbRpz = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRpz2729(in["slb_rpz"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemCtrLibAcct2480(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemCtrLibAcct2480 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemCtrLibAcct2480
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsInc2481(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsRate2482(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsInc2481(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsInc2481 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsInc2481
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TotalNodesFreeFailed = in["total_nodes_free_failed"].(int)
		ret.TotalNodesUnlinkFailed = in["total_nodes_unlink_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsRate2482(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsRate2482 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsRate2482
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.TotalNodesFreeFailed = in["total_nodes_free_failed"].(int)
		ret.TotalNodesUnlinkFailed = in["total_nodes_unlink_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerate2483(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerate2483 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerate2483
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsInc2484(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsRate2485(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsInc2484(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsInc2484 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsInc2484
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HwFwdProgErrors = in["hw_fwd_prog_errors"].(int)
		ret.HwFwdFlowSinglebitErrors = in["hw_fwd_flow_singlebit_errors"].(int)
		ret.HwFwdFlowTagMismatch = in["hw_fwd_flow_tag_mismatch"].(int)
		ret.HwFwdFlowSeqMismatch = in["hw_fwd_flow_seq_mismatch"].(int)
		ret.HwFwdFlowErrorCount = in["hw_fwd_flow_error_count"].(int)
		ret.HwFwdFlowUnalignCount = in["hw_fwd_flow_unalign_count"].(int)
		ret.HwFwdFlowUnderflowCount = in["hw_fwd_flow_underflow_count"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsRate2485(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsRate2485 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsRate2485
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.HwFwdProgErrors = in["hw_fwd_prog_errors"].(int)
		ret.HwFwdFlowSinglebitErrors = in["hw_fwd_flow_singlebit_errors"].(int)
		ret.HwFwdFlowTagMismatch = in["hw_fwd_flow_tag_mismatch"].(int)
		ret.HwFwdFlowSeqMismatch = in["hw_fwd_flow_seq_mismatch"].(int)
		ret.HwFwdFlowErrorCount = in["hw_fwd_flow_error_count"].(int)
		ret.HwFwdFlowUnalignCount = in["hw_fwd_flow_unalign_count"].(int)
		ret.HwFwdFlowUnderflowCount = in["hw_fwd_flow_underflow_count"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemRadiusServer2486(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemRadiusServer2486 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemRadiusServer2486
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsInc2487(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsRate2488(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsInc2487(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsInc2487 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsInc2487
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RadiusRequestDropped = in["radius_request_dropped"].(int)
		ret.RequestBadSecretDropped = in["request_bad_secret_dropped"].(int)
		ret.RequestNoKeyVapDropped = in["request_no_key_vap_dropped"].(int)
		ret.RequestMalformedDropped = in["request_malformed_dropped"].(int)
		ret.RadiusTableFull = in["radius_table_full"].(int)
		ret.SecretNotConfiguredDropped = in["secret_not_configured_dropped"].(int)
		ret.HaStandbyDropped = in["ha_standby_dropped"].(int)
		ret.Ipv6PrefixLengthMismatch = in["ipv6_prefix_length_mismatch"].(int)
		ret.InvalidKey = in["invalid_key"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsRate2488(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsRate2488 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsRate2488
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.RadiusRequestDropped = in["radius_request_dropped"].(int)
		ret.RequestBadSecretDropped = in["request_bad_secret_dropped"].(int)
		ret.RequestNoKeyVapDropped = in["request_no_key_vap_dropped"].(int)
		ret.RequestMalformedDropped = in["request_malformed_dropped"].(int)
		ret.RadiusTableFull = in["radius_table_full"].(int)
		ret.SecretNotConfiguredDropped = in["secret_not_configured_dropped"].(int)
		ret.HaStandbyDropped = in["ha_standby_dropped"].(int)
		ret.Ipv6PrefixLengthMismatch = in["ipv6_prefix_length_mismatch"].(int)
		ret.InvalidKey = in["invalid_key"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemIpThreatList2489(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemIpThreatList2489 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemIpThreatList2489
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsInc2490(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsRate2491(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsInc2490(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsInc2490 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsInc2490
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error_out_of_memory = in["error_out_of_memory"].(int)
		ret.Error_out_of_spe_entries = in["error_out_of_spe_entries"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsRate2491(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsRate2491 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsRate2491
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Error_out_of_memory = in["error_out_of_memory"].(int)
		ret.Error_out_of_spe_entries = in["error_out_of_spe_entries"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop2492(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop2492 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop2492
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc2493(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsRate2494(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc2493(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc2493 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc2493
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MrxDrop = in["mrx_drop"].(int)
		ret.HrxDrop = in["hrx_drop"].(int)
		ret.SizDrop = in["siz_drop"].(int)
		ret.FcsDrop = in["fcs_drop"].(int)
		ret.LandDrop = in["land_drop"].(int)
		ret.EmptyFragDrop = in["empty_frag_drop"].(int)
		ret.MicFragDrop = in["mic_frag_drop"].(int)
		ret.Ipv4OptDrop = in["ipv4_opt_drop"].(int)
		ret.Ipv4Frag = in["ipv4_frag"].(int)
		ret.BadIpHdrLen = in["bad_ip_hdr_len"].(int)
		ret.BadIpFlagsDrop = in["bad_ip_flags_drop"].(int)
		ret.BadIpTtlDrop = in["bad_ip_ttl_drop"].(int)
		ret.NoIpPayloadDrop = in["no_ip_payload_drop"].(int)
		ret.OversizeIpPayload = in["oversize_ip_payload"].(int)
		ret.BadIpPayloadLen = in["bad_ip_payload_len"].(int)
		ret.BadIpFragOffset = in["bad_ip_frag_offset"].(int)
		ret.BadIpChksumDrop = in["bad_ip_chksum_drop"].(int)
		ret.IcmpPodDrop = in["icmp_pod_drop"].(int)
		ret.TcpBadUrgOffet = in["tcp_bad_urg_offet"].(int)
		ret.TcpShortHdr = in["tcp_short_hdr"].(int)
		ret.TcpBadIpLen = in["tcp_bad_ip_len"].(int)
		ret.TcpNullFlags = in["tcp_null_flags"].(int)
		ret.TcpNullScan = in["tcp_null_scan"].(int)
		ret.TcpFinSin = in["tcp_fin_sin"].(int)
		ret.TcpXmasFlags = in["tcp_xmas_flags"].(int)
		ret.TcpXmasScan = in["tcp_xmas_scan"].(int)
		ret.TcpSynFrag = in["tcp_syn_frag"].(int)
		ret.TcpFragHdr = in["tcp_frag_hdr"].(int)
		ret.TcpBadChksum = in["tcp_bad_chksum"].(int)
		ret.UdpShortHdr = in["udp_short_hdr"].(int)
		ret.UdpBadIpLen = in["udp_bad_ip_len"].(int)
		ret.UdpKbFrags = in["udp_kb_frags"].(int)
		ret.UdpPortLb = in["udp_port_lb"].(int)
		ret.UdpBadChksum = in["udp_bad_chksum"].(int)
		ret.RuntIpHdr = in["runt_ip_hdr"].(int)
		ret.RuntTcpudpHdr = in["runt_tcpudp_hdr"].(int)
		ret.TunMismatch = in["tun_mismatch"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsRate2494(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsRate2494 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsRate2494
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.MrxDrop = in["mrx_drop"].(int)
		ret.HrxDrop = in["hrx_drop"].(int)
		ret.SizDrop = in["siz_drop"].(int)
		ret.FcsDrop = in["fcs_drop"].(int)
		ret.LandDrop = in["land_drop"].(int)
		ret.EmptyFragDrop = in["empty_frag_drop"].(int)
		ret.MicFragDrop = in["mic_frag_drop"].(int)
		ret.Ipv4OptDrop = in["ipv4_opt_drop"].(int)
		ret.Ipv4Frag = in["ipv4_frag"].(int)
		ret.BadIpHdrLen = in["bad_ip_hdr_len"].(int)
		ret.BadIpFlagsDrop = in["bad_ip_flags_drop"].(int)
		ret.BadIpTtlDrop = in["bad_ip_ttl_drop"].(int)
		ret.NoIpPayloadDrop = in["no_ip_payload_drop"].(int)
		ret.OversizeIpPayload = in["oversize_ip_payload"].(int)
		ret.BadIpPayloadLen = in["bad_ip_payload_len"].(int)
		ret.BadIpFragOffset = in["bad_ip_frag_offset"].(int)
		ret.BadIpChksumDrop = in["bad_ip_chksum_drop"].(int)
		ret.IcmpPodDrop = in["icmp_pod_drop"].(int)
		ret.TcpBadUrgOffet = in["tcp_bad_urg_offet"].(int)
		ret.TcpShortHdr = in["tcp_short_hdr"].(int)
		ret.TcpBadIpLen = in["tcp_bad_ip_len"].(int)
		ret.TcpNullFlags = in["tcp_null_flags"].(int)
		ret.TcpNullScan = in["tcp_null_scan"].(int)
		ret.TcpFinSin = in["tcp_fin_sin"].(int)
		ret.TcpXmasFlags = in["tcp_xmas_flags"].(int)
		ret.TcpXmasScan = in["tcp_xmas_scan"].(int)
		ret.TcpSynFrag = in["tcp_syn_frag"].(int)
		ret.TcpFragHdr = in["tcp_frag_hdr"].(int)
		ret.TcpBadChksum = in["tcp_bad_chksum"].(int)
		ret.UdpShortHdr = in["udp_short_hdr"].(int)
		ret.UdpBadIpLen = in["udp_bad_ip_len"].(int)
		ret.UdpKbFrags = in["udp_kb_frags"].(int)
		ret.UdpPortLb = in["udp_port_lb"].(int)
		ret.UdpBadChksum = in["udp_bad_chksum"].(int)
		ret.RuntIpHdr = in["runt_ip_hdr"].(int)
		ret.RuntTcpudpHdr = in["runt_tcpudp_hdr"].(int)
		ret.TunMismatch = in["tun_mismatch"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats2495(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats2495 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats2495
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc2496(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate2497(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc2496(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc2496 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc2496
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PktDrop = in["pkt_drop"].(int)
		ret.PktLnkDownDrop = in["pkt_lnk_down_drop"].(int)
		ret.ErrPktDrop = in["err_pkt_drop"].(int)
		ret.RxErr = in["rx_err"].(int)
		ret.TxErr = in["tx_err"].(int)
		ret.TxDrop = in["tx_drop"].(int)
		ret.RxLenErr = in["rx_len_err"].(int)
		ret.RxOverErr = in["rx_over_err"].(int)
		ret.RxCrcErr = in["rx_crc_err"].(int)
		ret.RxFrameErr = in["rx_frame_err"].(int)
		ret.RxNoBuffErr = in["rx_no_buff_err"].(int)
		ret.RxMissErr = in["rx_miss_err"].(int)
		ret.TxAbortErr = in["tx_abort_err"].(int)
		ret.TxCarrierErr = in["tx_carrier_err"].(int)
		ret.TxFifoErr = in["tx_fifo_err"].(int)
		ret.TxHbeatErr = in["tx_hbeat_err"].(int)
		ret.TxWindowsErr = in["tx_windows_err"].(int)
		ret.RxLongLenErr = in["rx_long_len_err"].(int)
		ret.RxShortLenErr = in["rx_short_len_err"].(int)
		ret.RxAlignErr = in["rx_align_err"].(int)
		ret.RxCsumOffloadErr = in["rx_csum_offload_err"].(int)
		ret.IoRxQueDrop = in["io_rx_que_drop"].(int)
		ret.IoTxQueDrop = in["io_tx_que_drop"].(int)
		ret.IoRingDrop = in["io_ring_drop"].(int)
		ret.WTxQueDrop = in["w_tx_que_drop"].(int)
		ret.WLinkDownDrop = in["w_link_down_drop"].(int)
		ret.WRingDrop = in["w_ring_drop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate2497(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate2497 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate2497
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.PktDrop = in["pkt_drop"].(int)
		ret.PktLnkDownDrop = in["pkt_lnk_down_drop"].(int)
		ret.ErrPktDrop = in["err_pkt_drop"].(int)
		ret.RxErr = in["rx_err"].(int)
		ret.TxErr = in["tx_err"].(int)
		ret.TxDrop = in["tx_drop"].(int)
		ret.RxLenErr = in["rx_len_err"].(int)
		ret.RxOverErr = in["rx_over_err"].(int)
		ret.RxCrcErr = in["rx_crc_err"].(int)
		ret.RxFrameErr = in["rx_frame_err"].(int)
		ret.RxNoBuffErr = in["rx_no_buff_err"].(int)
		ret.RxMissErr = in["rx_miss_err"].(int)
		ret.TxAbortErr = in["tx_abort_err"].(int)
		ret.TxCarrierErr = in["tx_carrier_err"].(int)
		ret.TxFifoErr = in["tx_fifo_err"].(int)
		ret.TxHbeatErr = in["tx_hbeat_err"].(int)
		ret.TxWindowsErr = in["tx_windows_err"].(int)
		ret.RxLongLenErr = in["rx_long_len_err"].(int)
		ret.RxShortLenErr = in["rx_short_len_err"].(int)
		ret.RxAlignErr = in["rx_align_err"].(int)
		ret.RxCsumOffloadErr = in["rx_csum_offload_err"].(int)
		ret.IoRxQueDrop = in["io_rx_que_drop"].(int)
		ret.IoTxQueDrop = in["io_tx_que_drop"].(int)
		ret.IoRingDrop = in["io_ring_drop"].(int)
		ret.WTxQueDrop = in["w_tx_que_drop"].(int)
		ret.WLinkDownDrop = in["w_link_down_drop"].(int)
		ret.WRingDrop = in["w_ring_drop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDrop2498(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDrop2498 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDrop2498
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc2499(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsRate2500(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc2499(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc2499 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc2499
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Land = in["land"].(int)
		ret.Emp_frg = in["emp_frg"].(int)
		ret.Emp_mic_frg = in["emp_mic_frg"].(int)
		ret.Opt = in["opt"].(int)
		ret.Frg = in["frg"].(int)
		ret.Bad_ip_hdrlen = in["bad_ip_hdrlen"].(int)
		ret.Bad_ip_flg = in["bad_ip_flg"].(int)
		ret.Bad_ip_ttl = in["bad_ip_ttl"].(int)
		ret.No_ip_payload = in["no_ip_payload"].(int)
		ret.Over_ip_payload = in["over_ip_payload"].(int)
		ret.Bad_ip_payload_len = in["bad_ip_payload_len"].(int)
		ret.Bad_ip_frg_offset = in["bad_ip_frg_offset"].(int)
		ret.Csum = in["csum"].(int)
		ret.Pod = in["pod"].(int)
		ret.Bad_tcp_urg_offset = in["bad_tcp_urg_offset"].(int)
		ret.Tcp_sht_hdr = in["tcp_sht_hdr"].(int)
		ret.Tcp_bad_iplen = in["tcp_bad_iplen"].(int)
		ret.Tcp_null_frg = in["tcp_null_frg"].(int)
		ret.Tcp_null_scan = in["tcp_null_scan"].(int)
		ret.Tcp_syn_fin = in["tcp_syn_fin"].(int)
		ret.Tcp_xmas = in["tcp_xmas"].(int)
		ret.Tcp_xmas_scan = in["tcp_xmas_scan"].(int)
		ret.Tcp_syn_frg = in["tcp_syn_frg"].(int)
		ret.Tcp_frg_hdr = in["tcp_frg_hdr"].(int)
		ret.Tcp_bad_csum = in["tcp_bad_csum"].(int)
		ret.Udp_srt_hdr = in["udp_srt_hdr"].(int)
		ret.Udp_bad_len = in["udp_bad_len"].(int)
		ret.Udp_kerb_frg = in["udp_kerb_frg"].(int)
		ret.Udp_port_lb = in["udp_port_lb"].(int)
		ret.Udp_bad_csum = in["udp_bad_csum"].(int)
		ret.Runt_ip_hdr = in["runt_ip_hdr"].(int)
		ret.Runt_tcp_udp_hdr = in["runt_tcp_udp_hdr"].(int)
		ret.Ipip_tnl_msmtch = in["ipip_tnl_msmtch"].(int)
		ret.Tcp_opt_err = in["tcp_opt_err"].(int)
		ret.Ipip_tnl_err = in["ipip_tnl_err"].(int)
		ret.Vxlan_err = in["vxlan_err"].(int)
		ret.Nvgre_err = in["nvgre_err"].(int)
		ret.Gre_pptp_err = in["gre_pptp_err"].(int)
		ret.Ipv6_eh_hbh = in["ipv6_eh_hbh"].(int)
		ret.Ipv6_eh_dest = in["ipv6_eh_dest"].(int)
		ret.Ipv6_eh_routing = in["ipv6_eh_routing"].(int)
		ret.Ipv6_eh_frag = in["ipv6_eh_frag"].(int)
		ret.Ipv6_eh_ah = in["ipv6_eh_ah"].(int)
		ret.Ipv6_eh_esp = in["ipv6_eh_esp"].(int)
		ret.Ipv6_eh_mobility = in["ipv6_eh_mobility"].(int)
		ret.Ipv6_eh_none = in["ipv6_eh_none"].(int)
		ret.Ipv6_eh_other = in["ipv6_eh_other"].(int)
		ret.Ipv6_eh_malformed = in["ipv6_eh_malformed"].(int)
		ret.Tcp_udp_zero_port = in["tcp_udp_zero_port"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsRate2500(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsRate2500 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsRate2500
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Land = in["land"].(int)
		ret.Emp_frg = in["emp_frg"].(int)
		ret.Emp_mic_frg = in["emp_mic_frg"].(int)
		ret.Opt = in["opt"].(int)
		ret.Frg = in["frg"].(int)
		ret.Bad_ip_hdrlen = in["bad_ip_hdrlen"].(int)
		ret.Bad_ip_flg = in["bad_ip_flg"].(int)
		ret.Bad_ip_ttl = in["bad_ip_ttl"].(int)
		ret.No_ip_payload = in["no_ip_payload"].(int)
		ret.Over_ip_payload = in["over_ip_payload"].(int)
		ret.Bad_ip_payload_len = in["bad_ip_payload_len"].(int)
		ret.Bad_ip_frg_offset = in["bad_ip_frg_offset"].(int)
		ret.Csum = in["csum"].(int)
		ret.Pod = in["pod"].(int)
		ret.Bad_tcp_urg_offset = in["bad_tcp_urg_offset"].(int)
		ret.Tcp_sht_hdr = in["tcp_sht_hdr"].(int)
		ret.Tcp_bad_iplen = in["tcp_bad_iplen"].(int)
		ret.Tcp_null_frg = in["tcp_null_frg"].(int)
		ret.Tcp_null_scan = in["tcp_null_scan"].(int)
		ret.Tcp_syn_fin = in["tcp_syn_fin"].(int)
		ret.Tcp_xmas = in["tcp_xmas"].(int)
		ret.Tcp_xmas_scan = in["tcp_xmas_scan"].(int)
		ret.Tcp_syn_frg = in["tcp_syn_frg"].(int)
		ret.Tcp_frg_hdr = in["tcp_frg_hdr"].(int)
		ret.Tcp_bad_csum = in["tcp_bad_csum"].(int)
		ret.Udp_srt_hdr = in["udp_srt_hdr"].(int)
		ret.Udp_bad_len = in["udp_bad_len"].(int)
		ret.Udp_kerb_frg = in["udp_kerb_frg"].(int)
		ret.Udp_port_lb = in["udp_port_lb"].(int)
		ret.Udp_bad_csum = in["udp_bad_csum"].(int)
		ret.Runt_ip_hdr = in["runt_ip_hdr"].(int)
		ret.Runt_tcp_udp_hdr = in["runt_tcp_udp_hdr"].(int)
		ret.Ipip_tnl_msmtch = in["ipip_tnl_msmtch"].(int)
		ret.Tcp_opt_err = in["tcp_opt_err"].(int)
		ret.Ipip_tnl_err = in["ipip_tnl_err"].(int)
		ret.Vxlan_err = in["vxlan_err"].(int)
		ret.Nvgre_err = in["nvgre_err"].(int)
		ret.Gre_pptp_err = in["gre_pptp_err"].(int)
		ret.Ipv6_eh_hbh = in["ipv6_eh_hbh"].(int)
		ret.Ipv6_eh_dest = in["ipv6_eh_dest"].(int)
		ret.Ipv6_eh_routing = in["ipv6_eh_routing"].(int)
		ret.Ipv6_eh_frag = in["ipv6_eh_frag"].(int)
		ret.Ipv6_eh_ah = in["ipv6_eh_ah"].(int)
		ret.Ipv6_eh_esp = in["ipv6_eh_esp"].(int)
		ret.Ipv6_eh_mobility = in["ipv6_eh_mobility"].(int)
		ret.Ipv6_eh_none = in["ipv6_eh_none"].(int)
		ret.Ipv6_eh_other = in["ipv6_eh_other"].(int)
		ret.Ipv6_eh_malformed = in["ipv6_eh_malformed"].(int)
		ret.Tcp_udp_zero_port = in["tcp_udp_zero_port"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthenticationGlobal2501(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthenticationGlobal2501 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthenticationGlobal2501
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsInc2502(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsRate2503(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsInc2502(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsInc2502 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsInc2502
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Misses = in["misses"].(int)
		ret.OpenSocketFailed = in["open_socket_failed"].(int)
		ret.ConnectFailed = in["connect_failed"].(int)
		ret.CreateTimerFailed = in["create_timer_failed"].(int)
		ret.GetSocketOptionFailed = in["get_socket_option_failed"].(int)
		ret.AflexAuthzFail = in["aflex_authz_fail"].(int)
		ret.AuthnFailure = in["authn_failure"].(int)
		ret.AuthzFailure = in["authz_failure"].(int)
		ret.DnsResolveFailed = in["dns_resolve_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsRate2503(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsRate2503 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsRate2503
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Misses = in["misses"].(int)
		ret.OpenSocketFailed = in["open_socket_failed"].(int)
		ret.ConnectFailed = in["connect_failed"].(int)
		ret.CreateTimerFailed = in["create_timer_failed"].(int)
		ret.GetSocketOptionFailed = in["get_socket_option_failed"].(int)
		ret.AflexAuthzFail = in["aflex_authz_fail"].(int)
		ret.AuthnFailure = in["authn_failure"].(int)
		ret.AuthzFailure = in["authz_failure"].(int)
		ret.DnsResolveFailed = in["dns_resolve_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamRdns2504(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamRdns2504 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamRdns2504
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamRdnsTriggerStatsInc2505(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamRdnsTriggerStatsRate2506(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamRdnsTriggerStatsInc2505(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamRdnsTriggerStatsInc2505 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamRdnsTriggerStatsInc2505
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RequestDropped = in["request_dropped"].(int)
		ret.ResponseFailure = in["response_failure"].(int)
		ret.ResponseError = in["response_error"].(int)
		ret.ResponseTimeout = in["response_timeout"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamRdnsTriggerStatsRate2506(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamRdnsTriggerStatsRate2506 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamRdnsTriggerStatsRate2506
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.RequestDropped = in["request_dropped"].(int)
		ret.ResponseFailure = in["response_failure"].(int)
		ret.ResponseError = in["response_error"].(int)
		ret.ResponseTimeout = in["response_timeout"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerLdap2507(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerLdap2507 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerLdap2507
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsInc2508(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsRate2509(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsInc2508(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsInc2508 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsInc2508
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AdminBindFailure = in["admin_bind_failure"].(int)
		ret.BindFailure = in["bind_failure"].(int)
		ret.SearchFailure = in["search_failure"].(int)
		ret.AuthorizeFailure = in["authorize_failure"].(int)
		ret.TimeoutError = in["timeout_error"].(int)
		ret.OtherError = in["other_error"].(int)
		ret.RequestDropped = in["request_dropped"].(int)
		ret.ResponseFailure = in["response_failure"].(int)
		ret.ResponseError = in["response_error"].(int)
		ret.ResponseTimeout = in["response_timeout"].(int)
		ret.JobStartError = in["job_start_error"].(int)
		ret.PollingControlError = in["polling_control_error"].(int)
		ret.SslSessionFailure = in["ssl_session_failure"].(int)
		ret.PwChangeFailure = in["pw_change_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsRate2509(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsRate2509 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsRate2509
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.AdminBindFailure = in["admin_bind_failure"].(int)
		ret.BindFailure = in["bind_failure"].(int)
		ret.SearchFailure = in["search_failure"].(int)
		ret.AuthorizeFailure = in["authorize_failure"].(int)
		ret.TimeoutError = in["timeout_error"].(int)
		ret.OtherError = in["other_error"].(int)
		ret.RequestDropped = in["request_dropped"].(int)
		ret.ResponseFailure = in["response_failure"].(int)
		ret.ResponseError = in["response_error"].(int)
		ret.ResponseTimeout = in["response_timeout"].(int)
		ret.JobStartError = in["job_start_error"].(int)
		ret.PollingControlError = in["polling_control_error"].(int)
		ret.SslSessionFailure = in["ssl_session_failure"].(int)
		ret.PwChangeFailure = in["pw_change_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerOcsp2510(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerOcsp2510 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerOcsp2510
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsInc2511(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsRate2512(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsInc2511(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsInc2511 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsInc2511
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaplingRequestDropped = in["stapling_request_dropped"].(int)
		ret.StaplingResponseFailure = in["stapling_response_failure"].(int)
		ret.StaplingResponseError = in["stapling_response_error"].(int)
		ret.StaplingResponseTimeout = in["stapling_response_timeout"].(int)
		ret.RequestDropped = in["request_dropped"].(int)
		ret.ResponseFailure = in["response_failure"].(int)
		ret.ResponseError = in["response_error"].(int)
		ret.ResponseTimeout = in["response_timeout"].(int)
		ret.JobStartError = in["job_start_error"].(int)
		ret.PollingControlError = in["polling_control_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsRate2512(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsRate2512 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsRate2512
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.StaplingRequestDropped = in["stapling_request_dropped"].(int)
		ret.StaplingResponseFailure = in["stapling_response_failure"].(int)
		ret.StaplingResponseError = in["stapling_response_error"].(int)
		ret.StaplingResponseTimeout = in["stapling_response_timeout"].(int)
		ret.RequestDropped = in["request_dropped"].(int)
		ret.ResponseFailure = in["response_failure"].(int)
		ret.ResponseError = in["response_error"].(int)
		ret.ResponseTimeout = in["response_timeout"].(int)
		ret.JobStartError = in["job_start_error"].(int)
		ret.PollingControlError = in["polling_control_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerRadius2513(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerRadius2513 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerRadius2513
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsInc2514(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsRate2515(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsInc2514(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsInc2514 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsInc2514
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authen_failure = in["authen_failure"].(int)
		ret.Authorize_failure = in["authorize_failure"].(int)
		ret.Timeout_error = in["timeout_error"].(int)
		ret.Other_error = in["other_error"].(int)
		ret.RequestDropped = in["request_dropped"].(int)
		ret.ResponseFailure = in["response_failure"].(int)
		ret.ResponseError = in["response_error"].(int)
		ret.ResponseTimeout = in["response_timeout"].(int)
		ret.JobStartError = in["job_start_error"].(int)
		ret.PollingControlError = in["polling_control_error"].(int)
		ret.AccountingFailure = in["accounting_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsRate2515(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsRate2515 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsRate2515
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Authen_failure = in["authen_failure"].(int)
		ret.Authorize_failure = in["authorize_failure"].(int)
		ret.Timeout_error = in["timeout_error"].(int)
		ret.Other_error = in["other_error"].(int)
		ret.RequestDropped = in["request_dropped"].(int)
		ret.ResponseFailure = in["response_failure"].(int)
		ret.ResponseError = in["response_error"].(int)
		ret.ResponseTimeout = in["response_timeout"].(int)
		ret.JobStartError = in["job_start_error"].(int)
		ret.PollingControlError = in["polling_control_error"].(int)
		ret.AccountingFailure = in["accounting_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin2516(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin2516 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin2516
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc2517(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsRate2518(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc2517(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc2517 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc2517
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KerberosTimeoutError = in["kerberos_timeout_error"].(int)
		ret.KerberosOtherError = in["kerberos_other_error"].(int)
		ret.KerberosRequestDropped = in["kerberos_request_dropped"].(int)
		ret.KerberosResponseFailure = in["kerberos_response_failure"].(int)
		ret.KerberosResponseError = in["kerberos_response_error"].(int)
		ret.KerberosResponseTimeout = in["kerberos_response_timeout"].(int)
		ret.KerberosJobStartError = in["kerberos_job_start_error"].(int)
		ret.KerberosPollingControlError = in["kerberos_polling_control_error"].(int)
		ret.KerberosPwExpiry = in["kerberos_pw_expiry"].(int)
		ret.KerberosPwChangeFailure = in["kerberos_pw_change_failure"].(int)
		ret.KerberosValidateKdcFailure = in["kerberos_validate_kdc_failure"].(int)
		ret.KerberosGenerateKdcKeytabFailure = in["kerberos_generate_kdc_keytab_failure"].(int)
		ret.KerberosDeleteKdcKeytabFailure = in["kerberos_delete_kdc_keytab_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsRate2518(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsRate2518 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsRate2518
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.KerberosTimeoutError = in["kerberos_timeout_error"].(int)
		ret.KerberosOtherError = in["kerberos_other_error"].(int)
		ret.KerberosRequestDropped = in["kerberos_request_dropped"].(int)
		ret.KerberosResponseFailure = in["kerberos_response_failure"].(int)
		ret.KerberosResponseError = in["kerberos_response_error"].(int)
		ret.KerberosResponseTimeout = in["kerberos_response_timeout"].(int)
		ret.KerberosJobStartError = in["kerberos_job_start_error"].(int)
		ret.KerberosPollingControlError = in["kerberos_polling_control_error"].(int)
		ret.KerberosPwExpiry = in["kerberos_pw_expiry"].(int)
		ret.KerberosPwChangeFailure = in["kerberos_pw_change_failure"].(int)
		ret.KerberosValidateKdcFailure = in["kerberos_validate_kdc_failure"].(int)
		ret.KerberosGenerateKdcKeytabFailure = in["kerberos_generate_kdc_keytab_failure"].(int)
		ret.KerberosDeleteKdcKeytabFailure = in["kerberos_delete_kdc_keytab_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthAccount2519(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthAccount2519 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthAccount2519
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthAccountTriggerStatsInc2520(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthAccountTriggerStatsRate2521(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthAccountTriggerStatsInc2520(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthAccountTriggerStatsInc2520 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthAccountTriggerStatsInc2520
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RequestDropped = in["request_dropped"].(int)
		ret.ResponseFailure = in["response_failure"].(int)
		ret.ResponseError = in["response_error"].(int)
		ret.ResponseTimeout = in["response_timeout"].(int)
		ret.ResponseOther = in["response_other"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthAccountTriggerStatsRate2521(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthAccountTriggerStatsRate2521 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthAccountTriggerStatsRate2521
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.RequestDropped = in["request_dropped"].(int)
		ret.ResponseFailure = in["response_failure"].(int)
		ret.ResponseError = in["response_error"].(int)
		ret.ResponseTimeout = in["response_timeout"].(int)
		ret.ResponseOther = in["response_other"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthSamlGlobal2522(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthSamlGlobal2522 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthSamlGlobal2522
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsInc2523(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsRate2524(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsInc2523(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsInc2523 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsInc2523
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AcsAuthzFail = in["acs_authz_fail"].(int)
		ret.AcsError = in["acs_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsRate2524(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsRate2524 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsRate2524
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.AcsAuthzFail = in["acs_authz_fail"].(int)
		ret.AcsError = in["acs_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthRelayKerberos2525(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthRelayKerberos2525 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthRelayKerberos2525
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsInc2526(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsRate2527(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsInc2526(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsInc2526 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsInc2526
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TimeoutError = in["timeout_error"].(int)
		ret.OtherError = in["other_error"].(int)
		ret.RequestDropped = in["request_dropped"].(int)
		ret.ResponseFailure = in["response_failure"].(int)
		ret.ResponseError = in["response_error"].(int)
		ret.ResponseTimeout = in["response_timeout"].(int)
		ret.JobStartError = in["job_start_error"].(int)
		ret.PollingControlError = in["polling_control_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsRate2527(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsRate2527 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsRate2527
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.TimeoutError = in["timeout_error"].(int)
		ret.OtherError = in["other_error"].(int)
		ret.RequestDropped = in["request_dropped"].(int)
		ret.ResponseFailure = in["response_failure"].(int)
		ret.ResponseError = in["response_error"].(int)
		ret.ResponseTimeout = in["response_timeout"].(int)
		ret.JobStartError = in["job_start_error"].(int)
		ret.PollingControlError = in["polling_control_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthCaptcha2528(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthCaptcha2528 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthCaptcha2528
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsInc2529(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsRate2530(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsInc2529(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsInc2529 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsInc2529
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RequestDropped = in["request_dropped"].(int)
		ret.ResponseFailure = in["response_failure"].(int)
		ret.ResponseError = in["response_error"].(int)
		ret.ResponseTimeout = in["response_timeout"].(int)
		ret.JsonFail = in["json_fail"].(int)
		ret.AttrFail = in["attr_fail"].(int)
		ret.TimeoutError = in["timeout_error"].(int)
		ret.OtherError = in["other_error"].(int)
		ret.JobStartError = in["job_start_error"].(int)
		ret.PollingControlError = in["polling_control_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsRate2530(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsRate2530 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsRate2530
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.RequestDropped = in["request_dropped"].(int)
		ret.ResponseFailure = in["response_failure"].(int)
		ret.ResponseError = in["response_error"].(int)
		ret.ResponseTimeout = in["response_timeout"].(int)
		ret.JsonFail = in["json_fail"].(int)
		ret.AttrFail = in["attr_fail"].(int)
		ret.TimeoutError = in["timeout_error"].(int)
		ret.OtherError = in["other_error"].(int)
		ret.JobStartError = in["job_start_error"].(int)
		ret.PollingControlError = in["polling_control_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslError2531(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslError2531 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslError2531
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsInc2532(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate2533(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsInc2532(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsInc2532 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsInc2532
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AppDataInHandshake = in["app_data_in_handshake"].(int)
		ret.AttemptToReuseSessInDiffContext = in["attempt_to_reuse_sess_in_diff_context"].(int)
		ret.BadAlertRecord = in["bad_alert_record"].(int)
		ret.BadAuthenticationType = in["bad_authentication_type"].(int)
		ret.BadChangeCipherSpec = in["bad_change_cipher_spec"].(int)
		ret.BadChecksum = in["bad_checksum"].(int)
		ret.BadDataReturnedByCallback = in["bad_data_returned_by_callback"].(int)
		ret.BadDecompression = in["bad_decompression"].(int)
		ret.BadDhGLength = in["bad_dh_g_length"].(int)
		ret.BadDhPubKeyLength = in["bad_dh_pub_key_length"].(int)
		ret.BadDhPLength = in["bad_dh_p_length"].(int)
		ret.BadDigestLength = in["bad_digest_length"].(int)
		ret.BadDsaSignature = in["bad_dsa_signature"].(int)
		ret.BadHelloRequest = in["bad_hello_request"].(int)
		ret.BadLength = in["bad_length"].(int)
		ret.BadMacDecode = in["bad_mac_decode"].(int)
		ret.BadMessageType = in["bad_message_type"].(int)
		ret.BadPacketLength = in["bad_packet_length"].(int)
		ret.BadProtocolVersionCounter = in["bad_protocol_version_counter"].(int)
		ret.BadResponseArgument = in["bad_response_argument"].(int)
		ret.BadRsaDecrypt = in["bad_rsa_decrypt"].(int)
		ret.BadRsaEncrypt = in["bad_rsa_encrypt"].(int)
		ret.BadRsaELength = in["bad_rsa_e_length"].(int)
		ret.BadRsaModulusLength = in["bad_rsa_modulus_length"].(int)
		ret.BadRsaSignature = in["bad_rsa_signature"].(int)
		ret.BadSignature = in["bad_signature"].(int)
		ret.BadSslFiletype = in["bad_ssl_filetype"].(int)
		ret.BadSslSessionIdLength = in["bad_ssl_session_id_length"].(int)
		ret.BadState = in["bad_state"].(int)
		ret.BadWriteRetry = in["bad_write_retry"].(int)
		ret.BioNotSet = in["bio_not_set"].(int)
		ret.BlockCipherPadIsWrong = in["block_cipher_pad_is_wrong"].(int)
		ret.BnLib = in["bn_lib"].(int)
		ret.CaDnLengthMismatch = in["ca_dn_length_mismatch"].(int)
		ret.CaDnTooLong = in["ca_dn_too_long"].(int)
		ret.CcsReceivedEarly = in["ccs_received_early"].(int)
		ret.CertificateVerifyFailed = in["certificate_verify_failed"].(int)
		ret.CertLengthMismatch = in["cert_length_mismatch"].(int)
		ret.ChallengeIsDifferent = in["challenge_is_different"].(int)
		ret.CipherCodeWrongLength = in["cipher_code_wrong_length"].(int)
		ret.CipherOrHashUnavailable = in["cipher_or_hash_unavailable"].(int)
		ret.CipherTableSrcError = in["cipher_table_src_error"].(int)
		ret.CompressedLengthTooLong = in["compressed_length_too_long"].(int)
		ret.CompressionFailure = in["compression_failure"].(int)
		ret.CompressionLibraryError = in["compression_library_error"].(int)
		ret.ConnectionIdIsDifferent = in["connection_id_is_different"].(int)
		ret.ConnectionTypeNotSet = in["connection_type_not_set"].(int)
		ret.DataBetweenCcsAndFinished = in["data_between_ccs_and_finished"].(int)
		ret.DataLengthTooLong = in["data_length_too_long"].(int)
		ret.DecryptionFailed = in["decryption_failed"].(int)
		ret.DecryptionFailedOrBadRecordMac = in["decryption_failed_or_bad_record_mac"].(int)
		ret.DhPublicValueLengthIsWrong = in["dh_public_value_length_is_wrong"].(int)
		ret.DigestCheckFailed = in["digest_check_failed"].(int)
		ret.EncryptedLengthTooLong = in["encrypted_length_too_long"].(int)
		ret.ErrorGeneratingTmpRsaKey = in["error_generating_tmp_rsa_key"].(int)
		ret.ErrorInReceivedCipherList = in["error_in_received_cipher_list"].(int)
		ret.ExcessiveMessageSize = in["excessive_message_size"].(int)
		ret.ExtraDataInMessage = in["extra_data_in_message"].(int)
		ret.GotAFinBeforeACcs = in["got_a_fin_before_a_ccs"].(int)
		ret.HttpsProxyRequest = in["https_proxy_request"].(int)
		ret.HttpRequest = in["http_request"].(int)
		ret.IllegalPadding = in["illegal_padding"].(int)
		ret.InappropriateFallback = in["inappropriate_fallback"].(int)
		ret.InvalidChallengeLength = in["invalid_challenge_length"].(int)
		ret.InvalidCommand = in["invalid_command"].(int)
		ret.InvalidPurpose = in["invalid_purpose"].(int)
		ret.InvalidStatusResponse = in["invalid_status_response"].(int)
		ret.InvalidTrust = in["invalid_trust"].(int)
		ret.KeyArgTooLong = in["key_arg_too_long"].(int)
		ret.Krb5 = in["krb5"].(int)
		ret.Krb5ClientCcPrincipal = in["krb5_client_cc_principal"].(int)
		ret.Krb5ClientGetCred = in["krb5_client_get_cred"].(int)
		ret.Krb5ClientInit = in["krb5_client_init"].(int)
		ret.Krb5ClientMkReq = in["krb5_client_mk_req"].(int)
		ret.Krb5ServerBadTicket = in["krb5_server_bad_ticket"].(int)
		ret.Krb5ServerInit = in["krb5_server_init"].(int)
		ret.Krb5ServerRdReq = in["krb5_server_rd_req"].(int)
		ret.Krb5ServerTktExpired = in["krb5_server_tkt_expired"].(int)
		ret.Krb5ServerTktNotYetValid = in["krb5_server_tkt_not_yet_valid"].(int)
		ret.Krb5ServerTktSkew = in["krb5_server_tkt_skew"].(int)
		ret.LengthMismatch = in["length_mismatch"].(int)
		ret.LengthTooShort = in["length_too_short"].(int)
		ret.LibraryBug = in["library_bug"].(int)
		ret.LibraryHasNoCiphers = in["library_has_no_ciphers"].(int)
		ret.MastKeyTooLong = in["mast_key_too_long"].(int)
		ret.MessageTooLong = in["message_too_long"].(int)
		ret.MissingDhDsaCert = in["missing_dh_dsa_cert"].(int)
		ret.MissingDhKey = in["missing_dh_key"].(int)
		ret.MissingDhRsaCert = in["missing_dh_rsa_cert"].(int)
		ret.MissingDsaSigningCert = in["missing_dsa_signing_cert"].(int)
		ret.MissingExportTmpDhKey = in["missing_export_tmp_dh_key"].(int)
		ret.MissingExportTmpRsaKey = in["missing_export_tmp_rsa_key"].(int)
		ret.MissingRsaCertificate = in["missing_rsa_certificate"].(int)
		ret.MissingRsaEncryptingCert = in["missing_rsa_encrypting_cert"].(int)
		ret.MissingRsaSigningCert = in["missing_rsa_signing_cert"].(int)
		ret.MissingTmpDhKey = in["missing_tmp_dh_key"].(int)
		ret.MissingTmpRsaKey = in["missing_tmp_rsa_key"].(int)
		ret.MissingTmpRsaPkey = in["missing_tmp_rsa_pkey"].(int)
		ret.MissingVerifyMessage = in["missing_verify_message"].(int)
		ret.NonSslv2InitialPacket = in["non_sslv2_initial_packet"].(int)
		ret.NoCertificatesReturned = in["no_certificates_returned"].(int)
		ret.NoCertificateAssigned = in["no_certificate_assigned"].(int)
		ret.NoCertificateReturned = in["no_certificate_returned"].(int)
		ret.NoCertificateSet = in["no_certificate_set"].(int)
		ret.NoCertificateSpecified = in["no_certificate_specified"].(int)
		ret.NoCiphersAvailable = in["no_ciphers_available"].(int)
		ret.NoCiphersPassed = in["no_ciphers_passed"].(int)
		ret.NoCiphersSpecified = in["no_ciphers_specified"].(int)
		ret.NoCipherList = in["no_cipher_list"].(int)
		ret.NoCipherMatch = in["no_cipher_match"].(int)
		ret.NoClientCertReceived = in["no_client_cert_received"].(int)
		ret.NoCompressionSpecified = in["no_compression_specified"].(int)
		ret.NoMethodSpecified = in["no_method_specified"].(int)
		ret.NoPrivatekey = in["no_privatekey"].(int)
		ret.NoPrivateKeyAssigned = in["no_private_key_assigned"].(int)
		ret.NoProtocolsAvailable = in["no_protocols_available"].(int)
		ret.NoPublickey = in["no_publickey"].(int)
		ret.NoSharedCipher = in["no_shared_cipher"].(int)
		ret.NoVerifyCallback = in["no_verify_callback"].(int)
		ret.NullSslCtx = in["null_ssl_ctx"].(int)
		ret.NullSslMethodPassed = in["null_ssl_method_passed"].(int)
		ret.OldSessionCipherNotReturned = in["old_session_cipher_not_returned"].(int)
		ret.PacketLengthTooLong = in["packet_length_too_long"].(int)
		ret.PathTooLong = in["path_too_long"].(int)
		ret.PeerDidNotReturnACertificate = in["peer_did_not_return_a_certificate"].(int)
		ret.PeerError = in["peer_error"].(int)
		ret.PeerErrorCertificate = in["peer_error_certificate"].(int)
		ret.PeerErrorNoCertificate = in["peer_error_no_certificate"].(int)
		ret.PeerErrorNoCipher = in["peer_error_no_cipher"].(int)
		ret.PeerErrorUnsupportedCertificateType = in["peer_error_unsupported_certificate_type"].(int)
		ret.PreMacLengthTooLong = in["pre_mac_length_too_long"].(int)
		ret.ProblemsMappingCipherFunctions = in["problems_mapping_cipher_functions"].(int)
		ret.ProtocolIsShutdown = in["protocol_is_shutdown"].(int)
		ret.PublicKeyEncryptError = in["public_key_encrypt_error"].(int)
		ret.PublicKeyIsNotRsa = in["public_key_is_not_rsa"].(int)
		ret.PublicKeyNotRsa = in["public_key_not_rsa"].(int)
		ret.ReadBioNotSet = in["read_bio_not_set"].(int)
		ret.ReadWrongPacketType = in["read_wrong_packet_type"].(int)
		ret.RecordLengthMismatch = in["record_length_mismatch"].(int)
		ret.RecordTooLarge = in["record_too_large"].(int)
		ret.RecordTooSmall = in["record_too_small"].(int)
		ret.RequiredCipherMissing = in["required_cipher_missing"].(int)
		ret.ReuseCertLengthNotZero = in["reuse_cert_length_not_zero"].(int)
		ret.ReuseCertTypeNotZero = in["reuse_cert_type_not_zero"].(int)
		ret.ReuseCipherListNotZero = in["reuse_cipher_list_not_zero"].(int)
		ret.ScsvReceivedWhenRenegotiating = in["scsv_received_when_renegotiating"].(int)
		ret.SessionIdContextUninitialized = in["session_id_context_uninitialized"].(int)
		ret.ShortRead = in["short_read"].(int)
		ret.SignatureForNonSigningCertificate = in["signature_for_non_signing_certificate"].(int)
		ret.Ssl23DoingSessionIdReuse = in["ssl23_doing_session_id_reuse"].(int)
		ret.Ssl2ConnectionIdTooLong = in["ssl2_connection_id_too_long"].(int)
		ret.Ssl3SessionIdTooLong = in["ssl3_session_id_too_long"].(int)
		ret.Ssl3SessionIdTooShort = in["ssl3_session_id_too_short"].(int)
		ret.Sslv3AlertBadCertificate = in["sslv3_alert_bad_certificate"].(int)
		ret.Sslv3AlertBadRecordMac = in["sslv3_alert_bad_record_mac"].(int)
		ret.Sslv3AlertCertificateExpired = in["sslv3_alert_certificate_expired"].(int)
		ret.Sslv3AlertCertificateRevoked = in["sslv3_alert_certificate_revoked"].(int)
		ret.Sslv3AlertCertificateUnknown = in["sslv3_alert_certificate_unknown"].(int)
		ret.Sslv3AlertDecompressionFailure = in["sslv3_alert_decompression_failure"].(int)
		ret.Sslv3AlertHandshakeFailure = in["sslv3_alert_handshake_failure"].(int)
		ret.Sslv3AlertIllegalParameter = in["sslv3_alert_illegal_parameter"].(int)
		ret.Sslv3AlertNoCertificate = in["sslv3_alert_no_certificate"].(int)
		ret.Sslv3AlertPeerErrorCert = in["sslv3_alert_peer_error_cert"].(int)
		ret.Sslv3AlertPeerErrorNoCert = in["sslv3_alert_peer_error_no_cert"].(int)
		ret.Sslv3AlertPeerErrorNoCipher = in["sslv3_alert_peer_error_no_cipher"].(int)
		ret.Sslv3AlertPeerErrorUnsuppCertType = in["sslv3_alert_peer_error_unsupp_cert_type"].(int)
		ret.Sslv3AlertUnexpectedMsg = in["sslv3_alert_unexpected_msg"].(int)
		ret.Sslv3AlertUnknownRemoteErrType = in["sslv3_alert_unknown_remote_err_type"].(int)
		ret.Sslv3AlertUnspportedCert = in["sslv3_alert_unspported_cert"].(int)
		ret.SslCtxHasNoDefaultSslVersion = in["ssl_ctx_has_no_default_ssl_version"].(int)
		ret.SslHandshakeFailure = in["ssl_handshake_failure"].(int)
		ret.SslLibraryHasNoCiphers = in["ssl_library_has_no_ciphers"].(int)
		ret.SslSessionIdCallbackFailed = in["ssl_session_id_callback_failed"].(int)
		ret.SslSessionIdConflict = in["ssl_session_id_conflict"].(int)
		ret.SslSessionIdContextTooLong = in["ssl_session_id_context_too_long"].(int)
		ret.SslSessionIdHasBadLength = in["ssl_session_id_has_bad_length"].(int)
		ret.SslSessionIdIsDifferent = in["ssl_session_id_is_different"].(int)
		ret.Tlsv1AlertAccessDenied = in["tlsv1_alert_access_denied"].(int)
		ret.Tlsv1AlertDecodeError = in["tlsv1_alert_decode_error"].(int)
		ret.Tlsv1AlertDecryptionFailed = in["tlsv1_alert_decryption_failed"].(int)
		ret.Tlsv1AlertDecryptError = in["tlsv1_alert_decrypt_error"].(int)
		ret.Tlsv1AlertExportRestriction = in["tlsv1_alert_export_restriction"].(int)
		ret.Tlsv1AlertInsufficientSecurity = in["tlsv1_alert_insufficient_security"].(int)
		ret.Tlsv1AlertInternalError = in["tlsv1_alert_internal_error"].(int)
		ret.Tlsv1AlertNoRenegotiation = in["tlsv1_alert_no_renegotiation"].(int)
		ret.Tlsv1AlertProtocolVersion = in["tlsv1_alert_protocol_version"].(int)
		ret.Tlsv1AlertRecordOverflow = in["tlsv1_alert_record_overflow"].(int)
		ret.Tlsv1AlertUnknownCa = in["tlsv1_alert_unknown_ca"].(int)
		ret.Tlsv1AlertUserCancelled = in["tlsv1_alert_user_cancelled"].(int)
		ret.TlsClientCertReqWithAnonCipher = in["tls_client_cert_req_with_anon_cipher"].(int)
		ret.TlsPeerDidNotRespondWithCertList = in["tls_peer_did_not_respond_with_cert_list"].(int)
		ret.TlsRsaEncryptedValueLengthIsWrong = in["tls_rsa_encrypted_value_length_is_wrong"].(int)
		ret.TriedToUseUnsupportedCipher = in["tried_to_use_unsupported_cipher"].(int)
		ret.UnableToDecodeDhCerts = in["unable_to_decode_dh_certs"].(int)
		ret.UnableToExtractPublicKey = in["unable_to_extract_public_key"].(int)
		ret.UnableToFindDhParameters = in["unable_to_find_dh_parameters"].(int)
		ret.UnableToFindPublicKeyParameters = in["unable_to_find_public_key_parameters"].(int)
		ret.UnableToFindSslMethod = in["unable_to_find_ssl_method"].(int)
		ret.UnableToLoadSsl2Md5Routines = in["unable_to_load_ssl2_md5_routines"].(int)
		ret.UnableToLoadSsl3Md5Routines = in["unable_to_load_ssl3_md5_routines"].(int)
		ret.UnableToLoadSsl3Sha1Routines = in["unable_to_load_ssl3_sha1_routines"].(int)
		ret.UnexpectedMessage = in["unexpected_message"].(int)
		ret.UnexpectedRecord = in["unexpected_record"].(int)
		ret.Uninitialized = in["uninitialized"].(int)
		ret.UnknownAlertType = in["unknown_alert_type"].(int)
		ret.UnknownCertificateType = in["unknown_certificate_type"].(int)
		ret.UnknownCipherReturned = in["unknown_cipher_returned"].(int)
		ret.UnknownCipherType = in["unknown_cipher_type"].(int)
		ret.UnknownKeyExchangeType = in["unknown_key_exchange_type"].(int)
		ret.UnknownPkeyType = in["unknown_pkey_type"].(int)
		ret.UnknownProtocol = in["unknown_protocol"].(int)
		ret.UnknownRemoteErrorType = in["unknown_remote_error_type"].(int)
		ret.UnknownSslVersion = in["unknown_ssl_version"].(int)
		ret.UnknownState = in["unknown_state"].(int)
		ret.UnsupportedCipher = in["unsupported_cipher"].(int)
		ret.UnsupportedCompressionAlgorithm = in["unsupported_compression_algorithm"].(int)
		ret.UnsupportedOption = in["unsupported_option"].(int)
		ret.UnsupportedProtocol = in["unsupported_protocol"].(int)
		ret.UnsupportedSslVersion = in["unsupported_ssl_version"].(int)
		ret.UnsupportedStatusType = in["unsupported_status_type"].(int)
		ret.WriteBioNotSet = in["write_bio_not_set"].(int)
		ret.WrongCipherReturned = in["wrong_cipher_returned"].(int)
		ret.WrongMessageType = in["wrong_message_type"].(int)
		ret.WrongCounterOfKeyBits = in["wrong_counter_of_key_bits"].(int)
		ret.WrongSignatureLength = in["wrong_signature_length"].(int)
		ret.WrongSignatureSize = in["wrong_signature_size"].(int)
		ret.WrongSslVersion = in["wrong_ssl_version"].(int)
		ret.WrongVersionCounter = in["wrong_version_counter"].(int)
		ret.X509Lib = in["x509_lib"].(int)
		ret.X509VerificationSetupProblems = in["x509_verification_setup_problems"].(int)
		ret.ClienthelloTlsext = in["clienthello_tlsext"].(int)
		ret.ParseTlsext = in["parse_tlsext"].(int)
		ret.ServerhelloTlsext = in["serverhello_tlsext"].(int)
		ret.Ssl3ExtInvalidServername = in["ssl3_ext_invalid_servername"].(int)
		ret.Ssl3ExtInvalidServernameType = in["ssl3_ext_invalid_servername_type"].(int)
		ret.MultipleSgcRestarts = in["multiple_sgc_restarts"].(int)
		ret.TlsInvalidEcpointformatList = in["tls_invalid_ecpointformat_list"].(int)
		ret.BadEccCert = in["bad_ecc_cert"].(int)
		ret.BadEcdsaSig = in["bad_ecdsa_sig"].(int)
		ret.BadEcpoint = in["bad_ecpoint"].(int)
		ret.CookieMismatch = in["cookie_mismatch"].(int)
		ret.UnsupportedEllipticCurve = in["unsupported_elliptic_curve"].(int)
		ret.NoRequiredDigest = in["no_required_digest"].(int)
		ret.UnsupportedDigestType = in["unsupported_digest_type"].(int)
		ret.BadHandshakeLength = in["bad_handshake_length"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate2533(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate2533 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate2533
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.AppDataInHandshake = in["app_data_in_handshake"].(int)
		ret.AttemptToReuseSessInDiffContext = in["attempt_to_reuse_sess_in_diff_context"].(int)
		ret.BadAlertRecord = in["bad_alert_record"].(int)
		ret.BadAuthenticationType = in["bad_authentication_type"].(int)
		ret.BadChangeCipherSpec = in["bad_change_cipher_spec"].(int)
		ret.BadChecksum = in["bad_checksum"].(int)
		ret.BadDataReturnedByCallback = in["bad_data_returned_by_callback"].(int)
		ret.BadDecompression = in["bad_decompression"].(int)
		ret.BadDhGLength = in["bad_dh_g_length"].(int)
		ret.BadDhPubKeyLength = in["bad_dh_pub_key_length"].(int)
		ret.BadDhPLength = in["bad_dh_p_length"].(int)
		ret.BadDigestLength = in["bad_digest_length"].(int)
		ret.BadDsaSignature = in["bad_dsa_signature"].(int)
		ret.BadHelloRequest = in["bad_hello_request"].(int)
		ret.BadLength = in["bad_length"].(int)
		ret.BadMacDecode = in["bad_mac_decode"].(int)
		ret.BadMessageType = in["bad_message_type"].(int)
		ret.BadPacketLength = in["bad_packet_length"].(int)
		ret.BadProtocolVersionCounter = in["bad_protocol_version_counter"].(int)
		ret.BadResponseArgument = in["bad_response_argument"].(int)
		ret.BadRsaDecrypt = in["bad_rsa_decrypt"].(int)
		ret.BadRsaEncrypt = in["bad_rsa_encrypt"].(int)
		ret.BadRsaELength = in["bad_rsa_e_length"].(int)
		ret.BadRsaModulusLength = in["bad_rsa_modulus_length"].(int)
		ret.BadRsaSignature = in["bad_rsa_signature"].(int)
		ret.BadSignature = in["bad_signature"].(int)
		ret.BadSslFiletype = in["bad_ssl_filetype"].(int)
		ret.BadSslSessionIdLength = in["bad_ssl_session_id_length"].(int)
		ret.BadState = in["bad_state"].(int)
		ret.BadWriteRetry = in["bad_write_retry"].(int)
		ret.BioNotSet = in["bio_not_set"].(int)
		ret.BlockCipherPadIsWrong = in["block_cipher_pad_is_wrong"].(int)
		ret.BnLib = in["bn_lib"].(int)
		ret.CaDnLengthMismatch = in["ca_dn_length_mismatch"].(int)
		ret.CaDnTooLong = in["ca_dn_too_long"].(int)
		ret.CcsReceivedEarly = in["ccs_received_early"].(int)
		ret.CertificateVerifyFailed = in["certificate_verify_failed"].(int)
		ret.CertLengthMismatch = in["cert_length_mismatch"].(int)
		ret.ChallengeIsDifferent = in["challenge_is_different"].(int)
		ret.CipherCodeWrongLength = in["cipher_code_wrong_length"].(int)
		ret.CipherOrHashUnavailable = in["cipher_or_hash_unavailable"].(int)
		ret.CipherTableSrcError = in["cipher_table_src_error"].(int)
		ret.CompressedLengthTooLong = in["compressed_length_too_long"].(int)
		ret.CompressionFailure = in["compression_failure"].(int)
		ret.CompressionLibraryError = in["compression_library_error"].(int)
		ret.ConnectionIdIsDifferent = in["connection_id_is_different"].(int)
		ret.ConnectionTypeNotSet = in["connection_type_not_set"].(int)
		ret.DataBetweenCcsAndFinished = in["data_between_ccs_and_finished"].(int)
		ret.DataLengthTooLong = in["data_length_too_long"].(int)
		ret.DecryptionFailed = in["decryption_failed"].(int)
		ret.DecryptionFailedOrBadRecordMac = in["decryption_failed_or_bad_record_mac"].(int)
		ret.DhPublicValueLengthIsWrong = in["dh_public_value_length_is_wrong"].(int)
		ret.DigestCheckFailed = in["digest_check_failed"].(int)
		ret.EncryptedLengthTooLong = in["encrypted_length_too_long"].(int)
		ret.ErrorGeneratingTmpRsaKey = in["error_generating_tmp_rsa_key"].(int)
		ret.ErrorInReceivedCipherList = in["error_in_received_cipher_list"].(int)
		ret.ExcessiveMessageSize = in["excessive_message_size"].(int)
		ret.ExtraDataInMessage = in["extra_data_in_message"].(int)
		ret.GotAFinBeforeACcs = in["got_a_fin_before_a_ccs"].(int)
		ret.HttpsProxyRequest = in["https_proxy_request"].(int)
		ret.HttpRequest = in["http_request"].(int)
		ret.IllegalPadding = in["illegal_padding"].(int)
		ret.InappropriateFallback = in["inappropriate_fallback"].(int)
		ret.InvalidChallengeLength = in["invalid_challenge_length"].(int)
		ret.InvalidCommand = in["invalid_command"].(int)
		ret.InvalidPurpose = in["invalid_purpose"].(int)
		ret.InvalidStatusResponse = in["invalid_status_response"].(int)
		ret.InvalidTrust = in["invalid_trust"].(int)
		ret.KeyArgTooLong = in["key_arg_too_long"].(int)
		ret.Krb5 = in["krb5"].(int)
		ret.Krb5ClientCcPrincipal = in["krb5_client_cc_principal"].(int)
		ret.Krb5ClientGetCred = in["krb5_client_get_cred"].(int)
		ret.Krb5ClientInit = in["krb5_client_init"].(int)
		ret.Krb5ClientMkReq = in["krb5_client_mk_req"].(int)
		ret.Krb5ServerBadTicket = in["krb5_server_bad_ticket"].(int)
		ret.Krb5ServerInit = in["krb5_server_init"].(int)
		ret.Krb5ServerRdReq = in["krb5_server_rd_req"].(int)
		ret.Krb5ServerTktExpired = in["krb5_server_tkt_expired"].(int)
		ret.Krb5ServerTktNotYetValid = in["krb5_server_tkt_not_yet_valid"].(int)
		ret.Krb5ServerTktSkew = in["krb5_server_tkt_skew"].(int)
		ret.LengthMismatch = in["length_mismatch"].(int)
		ret.LengthTooShort = in["length_too_short"].(int)
		ret.LibraryBug = in["library_bug"].(int)
		ret.LibraryHasNoCiphers = in["library_has_no_ciphers"].(int)
		ret.MastKeyTooLong = in["mast_key_too_long"].(int)
		ret.MessageTooLong = in["message_too_long"].(int)
		ret.MissingDhDsaCert = in["missing_dh_dsa_cert"].(int)
		ret.MissingDhKey = in["missing_dh_key"].(int)
		ret.MissingDhRsaCert = in["missing_dh_rsa_cert"].(int)
		ret.MissingDsaSigningCert = in["missing_dsa_signing_cert"].(int)
		ret.MissingExportTmpDhKey = in["missing_export_tmp_dh_key"].(int)
		ret.MissingExportTmpRsaKey = in["missing_export_tmp_rsa_key"].(int)
		ret.MissingRsaCertificate = in["missing_rsa_certificate"].(int)
		ret.MissingRsaEncryptingCert = in["missing_rsa_encrypting_cert"].(int)
		ret.MissingRsaSigningCert = in["missing_rsa_signing_cert"].(int)
		ret.MissingTmpDhKey = in["missing_tmp_dh_key"].(int)
		ret.MissingTmpRsaKey = in["missing_tmp_rsa_key"].(int)
		ret.MissingTmpRsaPkey = in["missing_tmp_rsa_pkey"].(int)
		ret.MissingVerifyMessage = in["missing_verify_message"].(int)
		ret.NonSslv2InitialPacket = in["non_sslv2_initial_packet"].(int)
		ret.NoCertificatesReturned = in["no_certificates_returned"].(int)
		ret.NoCertificateAssigned = in["no_certificate_assigned"].(int)
		ret.NoCertificateReturned = in["no_certificate_returned"].(int)
		ret.NoCertificateSet = in["no_certificate_set"].(int)
		ret.NoCertificateSpecified = in["no_certificate_specified"].(int)
		ret.NoCiphersAvailable = in["no_ciphers_available"].(int)
		ret.NoCiphersPassed = in["no_ciphers_passed"].(int)
		ret.NoCiphersSpecified = in["no_ciphers_specified"].(int)
		ret.NoCipherList = in["no_cipher_list"].(int)
		ret.NoCipherMatch = in["no_cipher_match"].(int)
		ret.NoClientCertReceived = in["no_client_cert_received"].(int)
		ret.NoCompressionSpecified = in["no_compression_specified"].(int)
		ret.NoMethodSpecified = in["no_method_specified"].(int)
		ret.NoPrivatekey = in["no_privatekey"].(int)
		ret.NoPrivateKeyAssigned = in["no_private_key_assigned"].(int)
		ret.NoProtocolsAvailable = in["no_protocols_available"].(int)
		ret.NoPublickey = in["no_publickey"].(int)
		ret.NoSharedCipher = in["no_shared_cipher"].(int)
		ret.NoVerifyCallback = in["no_verify_callback"].(int)
		ret.NullSslCtx = in["null_ssl_ctx"].(int)
		ret.NullSslMethodPassed = in["null_ssl_method_passed"].(int)
		ret.OldSessionCipherNotReturned = in["old_session_cipher_not_returned"].(int)
		ret.PacketLengthTooLong = in["packet_length_too_long"].(int)
		ret.PathTooLong = in["path_too_long"].(int)
		ret.PeerDidNotReturnACertificate = in["peer_did_not_return_a_certificate"].(int)
		ret.PeerError = in["peer_error"].(int)
		ret.PeerErrorCertificate = in["peer_error_certificate"].(int)
		ret.PeerErrorNoCertificate = in["peer_error_no_certificate"].(int)
		ret.PeerErrorNoCipher = in["peer_error_no_cipher"].(int)
		ret.PeerErrorUnsupportedCertificateType = in["peer_error_unsupported_certificate_type"].(int)
		ret.PreMacLengthTooLong = in["pre_mac_length_too_long"].(int)
		ret.ProblemsMappingCipherFunctions = in["problems_mapping_cipher_functions"].(int)
		ret.ProtocolIsShutdown = in["protocol_is_shutdown"].(int)
		ret.PublicKeyEncryptError = in["public_key_encrypt_error"].(int)
		ret.PublicKeyIsNotRsa = in["public_key_is_not_rsa"].(int)
		ret.PublicKeyNotRsa = in["public_key_not_rsa"].(int)
		ret.ReadBioNotSet = in["read_bio_not_set"].(int)
		ret.ReadWrongPacketType = in["read_wrong_packet_type"].(int)
		ret.RecordLengthMismatch = in["record_length_mismatch"].(int)
		ret.RecordTooLarge = in["record_too_large"].(int)
		ret.RecordTooSmall = in["record_too_small"].(int)
		ret.RequiredCipherMissing = in["required_cipher_missing"].(int)
		ret.ReuseCertLengthNotZero = in["reuse_cert_length_not_zero"].(int)
		ret.ReuseCertTypeNotZero = in["reuse_cert_type_not_zero"].(int)
		ret.ReuseCipherListNotZero = in["reuse_cipher_list_not_zero"].(int)
		ret.ScsvReceivedWhenRenegotiating = in["scsv_received_when_renegotiating"].(int)
		ret.SessionIdContextUninitialized = in["session_id_context_uninitialized"].(int)
		ret.ShortRead = in["short_read"].(int)
		ret.SignatureForNonSigningCertificate = in["signature_for_non_signing_certificate"].(int)
		ret.Ssl23DoingSessionIdReuse = in["ssl23_doing_session_id_reuse"].(int)
		ret.Ssl2ConnectionIdTooLong = in["ssl2_connection_id_too_long"].(int)
		ret.Ssl3SessionIdTooLong = in["ssl3_session_id_too_long"].(int)
		ret.Ssl3SessionIdTooShort = in["ssl3_session_id_too_short"].(int)
		ret.Sslv3AlertBadCertificate = in["sslv3_alert_bad_certificate"].(int)
		ret.Sslv3AlertBadRecordMac = in["sslv3_alert_bad_record_mac"].(int)
		ret.Sslv3AlertCertificateExpired = in["sslv3_alert_certificate_expired"].(int)
		ret.Sslv3AlertCertificateRevoked = in["sslv3_alert_certificate_revoked"].(int)
		ret.Sslv3AlertCertificateUnknown = in["sslv3_alert_certificate_unknown"].(int)
		ret.Sslv3AlertDecompressionFailure = in["sslv3_alert_decompression_failure"].(int)
		ret.Sslv3AlertHandshakeFailure = in["sslv3_alert_handshake_failure"].(int)
		ret.Sslv3AlertIllegalParameter = in["sslv3_alert_illegal_parameter"].(int)
		ret.Sslv3AlertNoCertificate = in["sslv3_alert_no_certificate"].(int)
		ret.Sslv3AlertPeerErrorCert = in["sslv3_alert_peer_error_cert"].(int)
		ret.Sslv3AlertPeerErrorNoCert = in["sslv3_alert_peer_error_no_cert"].(int)
		ret.Sslv3AlertPeerErrorNoCipher = in["sslv3_alert_peer_error_no_cipher"].(int)
		ret.Sslv3AlertPeerErrorUnsuppCertType = in["sslv3_alert_peer_error_unsupp_cert_type"].(int)
		ret.Sslv3AlertUnexpectedMsg = in["sslv3_alert_unexpected_msg"].(int)
		ret.Sslv3AlertUnknownRemoteErrType = in["sslv3_alert_unknown_remote_err_type"].(int)
		ret.Sslv3AlertUnspportedCert = in["sslv3_alert_unspported_cert"].(int)
		ret.SslCtxHasNoDefaultSslVersion = in["ssl_ctx_has_no_default_ssl_version"].(int)
		ret.SslHandshakeFailure = in["ssl_handshake_failure"].(int)
		ret.SslLibraryHasNoCiphers = in["ssl_library_has_no_ciphers"].(int)
		ret.SslSessionIdCallbackFailed = in["ssl_session_id_callback_failed"].(int)
		ret.SslSessionIdConflict = in["ssl_session_id_conflict"].(int)
		ret.SslSessionIdContextTooLong = in["ssl_session_id_context_too_long"].(int)
		ret.SslSessionIdHasBadLength = in["ssl_session_id_has_bad_length"].(int)
		ret.SslSessionIdIsDifferent = in["ssl_session_id_is_different"].(int)
		ret.Tlsv1AlertAccessDenied = in["tlsv1_alert_access_denied"].(int)
		ret.Tlsv1AlertDecodeError = in["tlsv1_alert_decode_error"].(int)
		ret.Tlsv1AlertDecryptionFailed = in["tlsv1_alert_decryption_failed"].(int)
		ret.Tlsv1AlertDecryptError = in["tlsv1_alert_decrypt_error"].(int)
		ret.Tlsv1AlertExportRestriction = in["tlsv1_alert_export_restriction"].(int)
		ret.Tlsv1AlertInsufficientSecurity = in["tlsv1_alert_insufficient_security"].(int)
		ret.Tlsv1AlertInternalError = in["tlsv1_alert_internal_error"].(int)
		ret.Tlsv1AlertNoRenegotiation = in["tlsv1_alert_no_renegotiation"].(int)
		ret.Tlsv1AlertProtocolVersion = in["tlsv1_alert_protocol_version"].(int)
		ret.Tlsv1AlertRecordOverflow = in["tlsv1_alert_record_overflow"].(int)
		ret.Tlsv1AlertUnknownCa = in["tlsv1_alert_unknown_ca"].(int)
		ret.Tlsv1AlertUserCancelled = in["tlsv1_alert_user_cancelled"].(int)
		ret.TlsClientCertReqWithAnonCipher = in["tls_client_cert_req_with_anon_cipher"].(int)
		ret.TlsPeerDidNotRespondWithCertList = in["tls_peer_did_not_respond_with_cert_list"].(int)
		ret.TlsRsaEncryptedValueLengthIsWrong = in["tls_rsa_encrypted_value_length_is_wrong"].(int)
		ret.TriedToUseUnsupportedCipher = in["tried_to_use_unsupported_cipher"].(int)
		ret.UnableToDecodeDhCerts = in["unable_to_decode_dh_certs"].(int)
		ret.UnableToExtractPublicKey = in["unable_to_extract_public_key"].(int)
		ret.UnableToFindDhParameters = in["unable_to_find_dh_parameters"].(int)
		ret.UnableToFindPublicKeyParameters = in["unable_to_find_public_key_parameters"].(int)
		ret.UnableToFindSslMethod = in["unable_to_find_ssl_method"].(int)
		ret.UnableToLoadSsl2Md5Routines = in["unable_to_load_ssl2_md5_routines"].(int)
		ret.UnableToLoadSsl3Md5Routines = in["unable_to_load_ssl3_md5_routines"].(int)
		ret.UnableToLoadSsl3Sha1Routines = in["unable_to_load_ssl3_sha1_routines"].(int)
		ret.UnexpectedMessage = in["unexpected_message"].(int)
		ret.UnexpectedRecord = in["unexpected_record"].(int)
		ret.Uninitialized = in["uninitialized"].(int)
		ret.UnknownAlertType = in["unknown_alert_type"].(int)
		ret.UnknownCertificateType = in["unknown_certificate_type"].(int)
		ret.UnknownCipherReturned = in["unknown_cipher_returned"].(int)
		ret.UnknownCipherType = in["unknown_cipher_type"].(int)
		ret.UnknownKeyExchangeType = in["unknown_key_exchange_type"].(int)
		ret.UnknownPkeyType = in["unknown_pkey_type"].(int)
		ret.UnknownProtocol = in["unknown_protocol"].(int)
		ret.UnknownRemoteErrorType = in["unknown_remote_error_type"].(int)
		ret.UnknownSslVersion = in["unknown_ssl_version"].(int)
		ret.UnknownState = in["unknown_state"].(int)
		ret.UnsupportedCipher = in["unsupported_cipher"].(int)
		ret.UnsupportedCompressionAlgorithm = in["unsupported_compression_algorithm"].(int)
		ret.UnsupportedOption = in["unsupported_option"].(int)
		ret.UnsupportedProtocol = in["unsupported_protocol"].(int)
		ret.UnsupportedSslVersion = in["unsupported_ssl_version"].(int)
		ret.UnsupportedStatusType = in["unsupported_status_type"].(int)
		ret.WriteBioNotSet = in["write_bio_not_set"].(int)
		ret.WrongCipherReturned = in["wrong_cipher_returned"].(int)
		ret.WrongMessageType = in["wrong_message_type"].(int)
		ret.WrongCounterOfKeyBits = in["wrong_counter_of_key_bits"].(int)
		ret.WrongSignatureLength = in["wrong_signature_length"].(int)
		ret.WrongSignatureSize = in["wrong_signature_size"].(int)
		ret.WrongSslVersion = in["wrong_ssl_version"].(int)
		ret.WrongVersionCounter = in["wrong_version_counter"].(int)
		ret.X509Lib = in["x509_lib"].(int)
		ret.X509VerificationSetupProblems = in["x509_verification_setup_problems"].(int)
		ret.ClienthelloTlsext = in["clienthello_tlsext"].(int)
		ret.ParseTlsext = in["parse_tlsext"].(int)
		ret.ServerhelloTlsext = in["serverhello_tlsext"].(int)
		ret.Ssl3ExtInvalidServername = in["ssl3_ext_invalid_servername"].(int)
		ret.Ssl3ExtInvalidServernameType = in["ssl3_ext_invalid_servername_type"].(int)
		ret.MultipleSgcRestarts = in["multiple_sgc_restarts"].(int)
		ret.TlsInvalidEcpointformatList = in["tls_invalid_ecpointformat_list"].(int)
		ret.BadEccCert = in["bad_ecc_cert"].(int)
		ret.BadEcdsaSig = in["bad_ecdsa_sig"].(int)
		ret.BadEcpoint = in["bad_ecpoint"].(int)
		ret.CookieMismatch = in["cookie_mismatch"].(int)
		ret.UnsupportedEllipticCurve = in["unsupported_elliptic_curve"].(int)
		ret.NoRequiredDigest = in["no_required_digest"].(int)
		ret.UnsupportedDigestType = in["unsupported_digest_type"].(int)
		ret.BadHandshakeLength = in["bad_handshake_length"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke2534(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke2534 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke2534
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsInc2535(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate2536(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsInc2535(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsInc2535 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsInc2535
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ocsp_chain_status_revoked = in["ocsp_chain_status_revoked"].(int)
		ret.Ocsp_chain_status_unknown = in["ocsp_chain_status_unknown"].(int)
		ret.Ocsp_connection_error = in["ocsp_connection_error"].(int)
		ret.Ocsp_uri_not_found = in["ocsp_uri_not_found"].(int)
		ret.Ocsp_uri_https = in["ocsp_uri_https"].(int)
		ret.Ocsp_uri_unsupported = in["ocsp_uri_unsupported"].(int)
		ret.Ocsp_response_status_revoked = in["ocsp_response_status_revoked"].(int)
		ret.Ocsp_response_status_unknown = in["ocsp_response_status_unknown"].(int)
		ret.Ocsp_cache_status_revoked = in["ocsp_cache_status_revoked"].(int)
		ret.Ocsp_cache_miss = in["ocsp_cache_miss"].(int)
		ret.Ocsp_other_error = in["ocsp_other_error"].(int)
		ret.Ocsp_response_no_nonce = in["ocsp_response_no_nonce"].(int)
		ret.Ocsp_response_nonce_error = in["ocsp_response_nonce_error"].(int)
		ret.Crl_connection_error = in["crl_connection_error"].(int)
		ret.Crl_uri_not_found = in["crl_uri_not_found"].(int)
		ret.Crl_uri_https = in["crl_uri_https"].(int)
		ret.Crl_uri_unsupported = in["crl_uri_unsupported"].(int)
		ret.Crl_response_status_revoked = in["crl_response_status_revoked"].(int)
		ret.Crl_response_status_unknown = in["crl_response_status_unknown"].(int)
		ret.Crl_cache_status_revoked = in["crl_cache_status_revoked"].(int)
		ret.Crl_other_error = in["crl_other_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate2536(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate2536 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate2536
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Ocsp_chain_status_revoked = in["ocsp_chain_status_revoked"].(int)
		ret.Ocsp_chain_status_unknown = in["ocsp_chain_status_unknown"].(int)
		ret.Ocsp_connection_error = in["ocsp_connection_error"].(int)
		ret.Ocsp_uri_not_found = in["ocsp_uri_not_found"].(int)
		ret.Ocsp_uri_https = in["ocsp_uri_https"].(int)
		ret.Ocsp_uri_unsupported = in["ocsp_uri_unsupported"].(int)
		ret.Ocsp_response_status_revoked = in["ocsp_response_status_revoked"].(int)
		ret.Ocsp_response_status_unknown = in["ocsp_response_status_unknown"].(int)
		ret.Ocsp_cache_status_revoked = in["ocsp_cache_status_revoked"].(int)
		ret.Ocsp_cache_miss = in["ocsp_cache_miss"].(int)
		ret.Ocsp_other_error = in["ocsp_other_error"].(int)
		ret.Ocsp_response_no_nonce = in["ocsp_response_no_nonce"].(int)
		ret.Ocsp_response_nonce_error = in["ocsp_response_nonce_error"].(int)
		ret.Crl_connection_error = in["crl_connection_error"].(int)
		ret.Crl_uri_not_found = in["crl_uri_not_found"].(int)
		ret.Crl_uri_https = in["crl_uri_https"].(int)
		ret.Crl_uri_unsupported = in["crl_uri_unsupported"].(int)
		ret.Crl_response_status_revoked = in["crl_response_status_revoked"].(int)
		ret.Crl_response_status_unknown = in["crl_response_status_unknown"].(int)
		ret.Crl_cache_status_revoked = in["crl_cache_status_revoked"].(int)
		ret.Crl_other_error = in["crl_other_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxy2537(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxy2537 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxy2537
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsInc2538(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsRate2539(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsInc2538(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsInc2538 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsInc2538
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FailedInSslHandshakes = in["failed_in_ssl_handshakes"].(int)
		ret.FailedInCryptoOperations = in["failed_in_crypto_operations"].(int)
		ret.FailedInTcp = in["failed_in_tcp"].(int)
		ret.FailedInCertificateVerification = in["failed_in_certificate_verification"].(int)
		ret.FailedInCertificateSigning = in["failed_in_certificate_signing"].(int)
		ret.InvalidOcspStaplingResponse = in["invalid_ocsp_stapling_response"].(int)
		ret.RevokedOcspResponse = in["revoked_ocsp_response"].(int)
		ret.UnsupportedSslVersion = in["unsupported_ssl_version"].(int)
		ret.ConnectionsFailed = in["connections_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsRate2539(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsRate2539 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsRate2539
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.FailedInSslHandshakes = in["failed_in_ssl_handshakes"].(int)
		ret.FailedInCryptoOperations = in["failed_in_crypto_operations"].(int)
		ret.FailedInTcp = in["failed_in_tcp"].(int)
		ret.FailedInCertificateVerification = in["failed_in_certificate_verification"].(int)
		ret.FailedInCertificateSigning = in["failed_in_certificate_signing"].(int)
		ret.InvalidOcspStaplingResponse = in["invalid_ocsp_stapling_response"].(int)
		ret.RevokedOcspResponse = in["revoked_ocsp_response"].(int)
		ret.UnsupportedSslVersion = in["unsupported_ssl_version"].(int)
		ret.ConnectionsFailed = in["connections_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError2540(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError2540 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError2540
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsInc2541(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate2542(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsInc2541(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsInc2541 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsInc2541
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bad_opcode = in["bad_opcode"].(int)
		ret.Bad_sg_write_len = in["bad_sg_write_len"].(int)
		ret.Bad_len = in["bad_len"].(int)
		ret.Bad_ipsec_protocol = in["bad_ipsec_protocol"].(int)
		ret.Bad_ipsec_auth = in["bad_ipsec_auth"].(int)
		ret.Bad_ipsec_padding = in["bad_ipsec_padding"].(int)
		ret.Bad_ip_version = in["bad_ip_version"].(int)
		ret.Bad_auth_type = in["bad_auth_type"].(int)
		ret.Bad_encrypt_type = in["bad_encrypt_type"].(int)
		ret.Bad_ipsec_spi = in["bad_ipsec_spi"].(int)
		ret.Bad_checksum = in["bad_checksum"].(int)
		ret.Bad_ipsec_context = in["bad_ipsec_context"].(int)
		ret.Bad_ipsec_context_direction = in["bad_ipsec_context_direction"].(int)
		ret.Bad_ipsec_context_flag_mismatch = in["bad_ipsec_context_flag_mismatch"].(int)
		ret.Ipcomp_payload = in["ipcomp_payload"].(int)
		ret.Bad_selector_match = in["bad_selector_match"].(int)
		ret.Bad_fragment_size = in["bad_fragment_size"].(int)
		ret.Bad_inline_data = in["bad_inline_data"].(int)
		ret.Bad_frag_size_configuration = in["bad_frag_size_configuration"].(int)
		ret.Dummy_payload = in["dummy_payload"].(int)
		ret.Bad_ip_payload_type = in["bad_ip_payload_type"].(int)
		ret.Bad_min_frag_size_auth_sha384_512 = in["bad_min_frag_size_auth_sha384_512"].(int)
		ret.Bad_esp_next_header = in["bad_esp_next_header"].(int)
		ret.Bad_gre_header = in["bad_gre_header"].(int)
		ret.Bad_gre_protocol = in["bad_gre_protocol"].(int)
		ret.Ipv6_extension_headers_too_big = in["ipv6_extension_headers_too_big"].(int)
		ret.Ipv6_hop_by_hop_error = in["ipv6_hop_by_hop_error"].(int)
		ret.Error_ipv6_decrypt_rh_segs_left_error = in["error_ipv6_decrypt_rh_segs_left_error"].(int)
		ret.Ipv6_rh_length_error = in["ipv6_rh_length_error"].(int)
		ret.Ipv6_outbound_rh_copy_addr_error = in["ipv6_outbound_rh_copy_addr_error"].(int)
		ret.Error_ipv6_extension_header_bad = in["error_ipv6_extension_header_bad"].(int)
		ret.Bad_encrypt_type_ctr_gcm = in["bad_encrypt_type_ctr_gcm"].(int)
		ret.Ah_not_supported_with_gcm_gmac_sha2 = in["ah_not_supported_with_gcm_gmac_sha2"].(int)
		ret.Tfc_padding_with_prefrag_not_supported = in["tfc_padding_with_prefrag_not_supported"].(int)
		ret.Bad_srtp_auth_tag = in["bad_srtp_auth_tag"].(int)
		ret.Bad_ipcomp_configuration = in["bad_ipcomp_configuration"].(int)
		ret.Dsiv_incorrect_param = in["dsiv_incorrect_param"].(int)
		ret.Bad_ipsec_unknown = in["bad_ipsec_unknown"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate2542(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate2542 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate2542
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Bad_opcode = in["bad_opcode"].(int)
		ret.Bad_sg_write_len = in["bad_sg_write_len"].(int)
		ret.Bad_len = in["bad_len"].(int)
		ret.Bad_ipsec_protocol = in["bad_ipsec_protocol"].(int)
		ret.Bad_ipsec_auth = in["bad_ipsec_auth"].(int)
		ret.Bad_ipsec_padding = in["bad_ipsec_padding"].(int)
		ret.Bad_ip_version = in["bad_ip_version"].(int)
		ret.Bad_auth_type = in["bad_auth_type"].(int)
		ret.Bad_encrypt_type = in["bad_encrypt_type"].(int)
		ret.Bad_ipsec_spi = in["bad_ipsec_spi"].(int)
		ret.Bad_checksum = in["bad_checksum"].(int)
		ret.Bad_ipsec_context = in["bad_ipsec_context"].(int)
		ret.Bad_ipsec_context_direction = in["bad_ipsec_context_direction"].(int)
		ret.Bad_ipsec_context_flag_mismatch = in["bad_ipsec_context_flag_mismatch"].(int)
		ret.Ipcomp_payload = in["ipcomp_payload"].(int)
		ret.Bad_selector_match = in["bad_selector_match"].(int)
		ret.Bad_fragment_size = in["bad_fragment_size"].(int)
		ret.Bad_inline_data = in["bad_inline_data"].(int)
		ret.Bad_frag_size_configuration = in["bad_frag_size_configuration"].(int)
		ret.Dummy_payload = in["dummy_payload"].(int)
		ret.Bad_ip_payload_type = in["bad_ip_payload_type"].(int)
		ret.Bad_min_frag_size_auth_sha384_512 = in["bad_min_frag_size_auth_sha384_512"].(int)
		ret.Bad_esp_next_header = in["bad_esp_next_header"].(int)
		ret.Bad_gre_header = in["bad_gre_header"].(int)
		ret.Bad_gre_protocol = in["bad_gre_protocol"].(int)
		ret.Ipv6_extension_headers_too_big = in["ipv6_extension_headers_too_big"].(int)
		ret.Ipv6_hop_by_hop_error = in["ipv6_hop_by_hop_error"].(int)
		ret.Error_ipv6_decrypt_rh_segs_left_error = in["error_ipv6_decrypt_rh_segs_left_error"].(int)
		ret.Ipv6_rh_length_error = in["ipv6_rh_length_error"].(int)
		ret.Ipv6_outbound_rh_copy_addr_error = in["ipv6_outbound_rh_copy_addr_error"].(int)
		ret.Error_ipv6_extension_header_bad = in["error_ipv6_extension_header_bad"].(int)
		ret.Bad_encrypt_type_ctr_gcm = in["bad_encrypt_type_ctr_gcm"].(int)
		ret.Ah_not_supported_with_gcm_gmac_sha2 = in["ah_not_supported_with_gcm_gmac_sha2"].(int)
		ret.Tfc_padding_with_prefrag_not_supported = in["tfc_padding_with_prefrag_not_supported"].(int)
		ret.Bad_srtp_auth_tag = in["bad_srtp_auth_tag"].(int)
		ret.Bad_ipcomp_configuration = in["bad_ipcomp_configuration"].(int)
		ret.Dsiv_incorrect_param = in["dsiv_incorrect_param"].(int)
		ret.Bad_ipsec_unknown = in["bad_ipsec_unknown"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProc2543(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProc2543 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProc2543
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsInc2544(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsRate2545(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsInc2544(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsInc2544 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsInc2544
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L3_entry_match_drop = in["l3_entry_match_drop"].(int)
		ret.L3_entry_match_drop_hw = in["l3_entry_match_drop_hw"].(int)
		ret.L3_entry_drop_max_hw_exceeded = in["l3_entry_drop_max_hw_exceeded"].(int)
		ret.L4_entry_match_drop = in["l4_entry_match_drop"].(int)
		ret.L4_entry_match_drop_hw = in["l4_entry_match_drop_hw"].(int)
		ret.L4_entry_drop_max_hw_exceeded = in["l4_entry_drop_max_hw_exceeded"].(int)
		ret.L4_entry_list_alloc_failure = in["l4_entry_list_alloc_failure"].(int)
		ret.Ip_node_alloc_failure = in["ip_node_alloc_failure"].(int)
		ret.Ip_port_block_alloc_failure = in["ip_port_block_alloc_failure"].(int)
		ret.Ip_other_block_alloc_failure = in["ip_other_block_alloc_failure"].(int)
		ret.L3_entry_add_to_bgp_failure = in["l3_entry_add_to_bgp_failure"].(int)
		ret.L3_entry_remove_from_bgp_failure = in["l3_entry_remove_from_bgp_failure"].(int)
		ret.L3_entry_add_to_hw_failure = in["l3_entry_add_to_hw_failure"].(int)
		ret.Syn_cookie_verification_failed = in["syn_cookie_verification_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsRate2545(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsRate2545 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsRate2545
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.L3_entry_match_drop = in["l3_entry_match_drop"].(int)
		ret.L3_entry_match_drop_hw = in["l3_entry_match_drop_hw"].(int)
		ret.L3_entry_drop_max_hw_exceeded = in["l3_entry_drop_max_hw_exceeded"].(int)
		ret.L4_entry_match_drop = in["l4_entry_match_drop"].(int)
		ret.L4_entry_match_drop_hw = in["l4_entry_match_drop_hw"].(int)
		ret.L4_entry_drop_max_hw_exceeded = in["l4_entry_drop_max_hw_exceeded"].(int)
		ret.L4_entry_list_alloc_failure = in["l4_entry_list_alloc_failure"].(int)
		ret.Ip_node_alloc_failure = in["ip_node_alloc_failure"].(int)
		ret.Ip_port_block_alloc_failure = in["ip_port_block_alloc_failure"].(int)
		ret.Ip_other_block_alloc_failure = in["ip_other_block_alloc_failure"].(int)
		ret.L3_entry_add_to_bgp_failure = in["l3_entry_add_to_bgp_failure"].(int)
		ret.L3_entry_remove_from_bgp_failure = in["l3_entry_remove_from_bgp_failure"].(int)
		ret.L3_entry_add_to_hw_failure = in["l3_entry_add_to_hw_failure"].(int)
		ret.Syn_cookie_verification_failed = in["syn_cookie_verification_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Lsn2546(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Lsn2546 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Lsn2546
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc2547(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsRate2548(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc2547(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc2547 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc2547
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.User_quota_failure = in["user_quota_failure"].(int)
		ret.Data_sesn_user_quota_exceeded = in["data_sesn_user_quota_exceeded"].(int)
		ret.Fullcone_failure = in["fullcone_failure"].(int)
		ret.Fullcone_self_hairpinning_drop = in["fullcone_self_hairpinning_drop"].(int)
		ret.Nat_pool_unusable = in["nat_pool_unusable"].(int)
		ret.Ha_nat_pool_unusable = in["ha_nat_pool_unusable"].(int)
		ret.Ha_nat_pool_batch_type_mismatch = in["ha_nat_pool_batch_type_mismatch"].(int)
		ret.Sip_alg_quota_inc_failure = in["sip_alg_quota_inc_failure"].(int)
		ret.Sip_alg_alloc_rtp_rtcp_port_failure = in["sip_alg_alloc_rtp_rtcp_port_failure"].(int)
		ret.Sip_alg_alloc_single_port_failure = in["sip_alg_alloc_single_port_failure"].(int)
		ret.Sip_alg_create_single_fullcone_failure = in["sip_alg_create_single_fullcone_failure"].(int)
		ret.Sip_alg_create_rtp_fullcone_failure = in["sip_alg_create_rtp_fullcone_failure"].(int)
		ret.Sip_alg_create_rtcp_fullcone_failure = in["sip_alg_create_rtcp_fullcone_failure"].(int)
		ret.H323_alg_alloc_single_port_failure = in["h323_alg_alloc_single_port_failure"].(int)
		ret.H323_alg_create_single_fullcone_failure = in["h323_alg_create_single_fullcone_failure"].(int)
		ret.H323_alg_create_rtp_fullcone_failure = in["h323_alg_create_rtp_fullcone_failure"].(int)
		ret.H323_alg_create_rtcp_fullcone_failure = in["h323_alg_create_rtcp_fullcone_failure"].(int)
		ret.Port_overloading_out_of_memory = in["port_overloading_out_of_memory"].(int)
		ret.Port_overloading_inc_overflow = in["port_overloading_inc_overflow"].(int)
		ret.Fullcone_ext_mem_alloc_failure = in["fullcone_ext_mem_alloc_failure"].(int)
		ret.Fullcone_ext_mem_alloc_init_faulure = in["fullcone_ext_mem_alloc_init_faulure"].(int)
		ret.Mgcp_alg_create_rtp_fullcone_failure = in["mgcp_alg_create_rtp_fullcone_failure"].(int)
		ret.Mgcp_alg_create_rtcp_fullcone_failure = in["mgcp_alg_create_rtcp_fullcone_failure"].(int)
		ret.Mgcp_alg_port_pair_alloc_from_quota_par = in["mgcp_alg_port_pair_alloc_from_quota_par"].(int)
		ret.User_quota_unusable_drop = in["user_quota_unusable_drop"].(int)
		ret.User_quota_unusable = in["user_quota_unusable"].(int)
		ret.Adc_port_allocation_failed = in["adc_port_allocation_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsRate2548(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsRate2548 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsRate2548
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.User_quota_failure = in["user_quota_failure"].(int)
		ret.Data_sesn_user_quota_exceeded = in["data_sesn_user_quota_exceeded"].(int)
		ret.Fullcone_failure = in["fullcone_failure"].(int)
		ret.Fullcone_self_hairpinning_drop = in["fullcone_self_hairpinning_drop"].(int)
		ret.Nat_pool_unusable = in["nat_pool_unusable"].(int)
		ret.Ha_nat_pool_unusable = in["ha_nat_pool_unusable"].(int)
		ret.Ha_nat_pool_batch_type_mismatch = in["ha_nat_pool_batch_type_mismatch"].(int)
		ret.Sip_alg_quota_inc_failure = in["sip_alg_quota_inc_failure"].(int)
		ret.Sip_alg_alloc_rtp_rtcp_port_failure = in["sip_alg_alloc_rtp_rtcp_port_failure"].(int)
		ret.Sip_alg_alloc_single_port_failure = in["sip_alg_alloc_single_port_failure"].(int)
		ret.Sip_alg_create_single_fullcone_failure = in["sip_alg_create_single_fullcone_failure"].(int)
		ret.Sip_alg_create_rtp_fullcone_failure = in["sip_alg_create_rtp_fullcone_failure"].(int)
		ret.Sip_alg_create_rtcp_fullcone_failure = in["sip_alg_create_rtcp_fullcone_failure"].(int)
		ret.H323_alg_alloc_single_port_failure = in["h323_alg_alloc_single_port_failure"].(int)
		ret.H323_alg_create_single_fullcone_failure = in["h323_alg_create_single_fullcone_failure"].(int)
		ret.H323_alg_create_rtp_fullcone_failure = in["h323_alg_create_rtp_fullcone_failure"].(int)
		ret.H323_alg_create_rtcp_fullcone_failure = in["h323_alg_create_rtcp_fullcone_failure"].(int)
		ret.Port_overloading_out_of_memory = in["port_overloading_out_of_memory"].(int)
		ret.Port_overloading_inc_overflow = in["port_overloading_inc_overflow"].(int)
		ret.Fullcone_ext_mem_alloc_failure = in["fullcone_ext_mem_alloc_failure"].(int)
		ret.Fullcone_ext_mem_alloc_init_faulure = in["fullcone_ext_mem_alloc_init_faulure"].(int)
		ret.Mgcp_alg_create_rtp_fullcone_failure = in["mgcp_alg_create_rtp_fullcone_failure"].(int)
		ret.Mgcp_alg_create_rtcp_fullcone_failure = in["mgcp_alg_create_rtcp_fullcone_failure"].(int)
		ret.Mgcp_alg_port_pair_alloc_from_quota_par = in["mgcp_alg_port_pair_alloc_from_quota_par"].(int)
		ret.User_quota_unusable_drop = in["user_quota_unusable_drop"].(int)
		ret.User_quota_unusable = in["user_quota_unusable"].(int)
		ret.Adc_port_allocation_failed = in["adc_port_allocation_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgEsp2549(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgEsp2549 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgEsp2549
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsInc2550(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsRate2551(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsInc2550(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsInc2550 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsInc2550
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NatIpConflict = in["nat_ip_conflict"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsRate2551(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsRate2551 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsRate2551
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.NatIpConflict = in["nat_ip_conflict"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgPptp2552(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgPptp2552 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgPptp2552
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsInc2553(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsRate2554(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsInc2553(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsInc2553 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsInc2553
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NoGreSessionMatch = in["no_gre_session_match"].(int)
		ret.CallReqPnsCallIdMismatch = in["call_req_pns_call_id_mismatch"].(int)
		ret.CallReplyPnsCallIdMismatch = in["call_reply_pns_call_id_mismatch"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsRate2554(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsRate2554 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsRate2554
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.NoGreSessionMatch = in["no_gre_session_match"].(int)
		ret.CallReqPnsCallIdMismatch = in["call_req_pns_call_id_mismatch"].(int)
		ret.CallReplyPnsCallIdMismatch = in["call_reply_pns_call_id_mismatch"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgRtsp2555(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgRtsp2555 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgRtsp2555
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsInc2556(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsRate2557(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsInc2556(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsInc2556 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsInc2556
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StreamCreationFailure = in["stream_creation_failure"].(int)
		ret.PortAllocationFailure = in["port_allocation_failure"].(int)
		ret.UnknownClientPortFromServer = in["unknown_client_port_from_server"].(int)
		ret.NoSessionMem = in["no_session_mem"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsRate2557(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsRate2557 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsRate2557
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.StreamCreationFailure = in["stream_creation_failure"].(int)
		ret.PortAllocationFailure = in["port_allocation_failure"].(int)
		ret.UnknownClientPortFromServer = in["unknown_client_port_from_server"].(int)
		ret.NoSessionMem = in["no_session_mem"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSip2558(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSip2558 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSip2558
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsInc2559(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsRate2560(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsInc2559(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsInc2559 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsInc2559
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MethodUnknown = in["method_unknown"].(int)
		ret.ParseError = in["parse_error"].(int)
		ret.TcpOutOfOrderDrop = in["tcp_out_of_order_drop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsRate2560(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsRate2560 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsRate2560
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.MethodUnknown = in["method_unknown"].(int)
		ret.ParseError = in["parse_error"].(int)
		ret.TcpOutOfOrderDrop = in["tcp_out_of_order_drop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgMgcp2561(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgMgcp2561 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgMgcp2561
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsInc2562(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsRate2563(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsInc2562(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsInc2562 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsInc2562
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ParseError = in["parse_error"].(int)
		ret.TcpOutOfOrderDrop = in["tcp_out_of_order_drop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsRate2563(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsRate2563 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsRate2563
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.ParseError = in["parse_error"].(int)
		ret.TcpOutOfOrderDrop = in["tcp_out_of_order_drop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgH3232564(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgH3232564 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgH3232564
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsInc2565(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsRate2566(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsInc2565(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsInc2565 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsInc2565
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ParseError = in["parse_error"].(int)
		ret.TcpOutOfOrderDrop = in["tcp_out_of_order_drop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsRate2566(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsRate2566 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsRate2566
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.ParseError = in["parse_error"].(int)
		ret.TcpOutOfOrderDrop = in["tcp_out_of_order_drop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRadius2567(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRadius2567 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRadius2567
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsInc2568(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsRate2569(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsInc2568(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsInc2568 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsInc2568
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RadiusRequestDropped = in["radius_request_dropped"].(int)
		ret.RequestBadSecretDropped = in["request_bad_secret_dropped"].(int)
		ret.RequestNoKeyVapDropped = in["request_no_key_vap_dropped"].(int)
		ret.RequestMalformedDropped = in["request_malformed_dropped"].(int)
		ret.RequestIgnored = in["request_ignored"].(int)
		ret.RadiusTableFull = in["radius_table_full"].(int)
		ret.SecretNotConfiguredDropped = in["secret_not_configured_dropped"].(int)
		ret.HaStandbyDropped = in["ha_standby_dropped"].(int)
		ret.InvalidKey = in["invalid_key"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsRate2569(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsRate2569 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsRate2569
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.RadiusRequestDropped = in["radius_request_dropped"].(int)
		ret.RequestBadSecretDropped = in["request_bad_secret_dropped"].(int)
		ret.RequestNoKeyVapDropped = in["request_no_key_vap_dropped"].(int)
		ret.RequestMalformedDropped = in["request_malformed_dropped"].(int)
		ret.RequestIgnored = in["request_ignored"].(int)
		ret.RadiusTableFull = in["radius_table_full"].(int)
		ret.SecretNotConfiguredDropped = in["secret_not_configured_dropped"].(int)
		ret.HaStandbyDropped = in["ha_standby_dropped"].(int)
		ret.InvalidKey = in["invalid_key"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64Global2570(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64Global2570 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64Global2570
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc2571(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsRate2572(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc2571(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc2571 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc2571
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.User_quota_failure = in["user_quota_failure"].(int)
		ret.Nat_port_unavailable_tcp = in["nat_port_unavailable_tcp"].(int)
		ret.Nat_port_unavailable_udp = in["nat_port_unavailable_udp"].(int)
		ret.Nat_port_unavailable_icmp = in["nat_port_unavailable_icmp"].(int)
		ret.New_user_resource_unavailable = in["new_user_resource_unavailable"].(int)
		ret.Fullcone_failure = in["fullcone_failure"].(int)
		ret.Fullcone_self_hairpinning_drop = in["fullcone_self_hairpinning_drop"].(int)
		ret.Eif_limit_exceeded = in["eif_limit_exceeded"].(int)
		ret.Nat_pool_unusable = in["nat_pool_unusable"].(int)
		ret.Ha_nat_pool_unusable = in["ha_nat_pool_unusable"].(int)
		ret.Ha_nat_pool_batch_type_mismatch = in["ha_nat_pool_batch_type_mismatch"].(int)
		ret.No_radius_profile_match = in["no_radius_profile_match"].(int)
		ret.No_class_list_match = in["no_class_list_match"].(int)
		ret.User_quota_unusable_drop = in["user_quota_unusable_drop"].(int)
		ret.User_quota_unusable = in["user_quota_unusable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsRate2572(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsRate2572 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsRate2572
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.User_quota_failure = in["user_quota_failure"].(int)
		ret.Nat_port_unavailable_tcp = in["nat_port_unavailable_tcp"].(int)
		ret.Nat_port_unavailable_udp = in["nat_port_unavailable_udp"].(int)
		ret.Nat_port_unavailable_icmp = in["nat_port_unavailable_icmp"].(int)
		ret.New_user_resource_unavailable = in["new_user_resource_unavailable"].(int)
		ret.Fullcone_failure = in["fullcone_failure"].(int)
		ret.Fullcone_self_hairpinning_drop = in["fullcone_self_hairpinning_drop"].(int)
		ret.Eif_limit_exceeded = in["eif_limit_exceeded"].(int)
		ret.Nat_pool_unusable = in["nat_pool_unusable"].(int)
		ret.Ha_nat_pool_unusable = in["ha_nat_pool_unusable"].(int)
		ret.Ha_nat_pool_batch_type_mismatch = in["ha_nat_pool_batch_type_mismatch"].(int)
		ret.No_radius_profile_match = in["no_radius_profile_match"].(int)
		ret.No_class_list_match = in["no_class_list_match"].(int)
		ret.User_quota_unusable_drop = in["user_quota_unusable_drop"].(int)
		ret.User_quota_unusable = in["user_quota_unusable"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DsLiteGlobal2573(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DsLiteGlobal2573 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DsLiteGlobal2573
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsInc2574(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsRate2575(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsInc2574(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsInc2574 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsInc2574
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.User_quota_failure = in["user_quota_failure"].(int)
		ret.Nat_port_unavailable_tcp = in["nat_port_unavailable_tcp"].(int)
		ret.Nat_port_unavailable_udp = in["nat_port_unavailable_udp"].(int)
		ret.Nat_port_unavailable_icmp = in["nat_port_unavailable_icmp"].(int)
		ret.Fullcone_failure = in["fullcone_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsRate2575(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsRate2575 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsRate2575
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.User_quota_failure = in["user_quota_failure"].(int)
		ret.Nat_port_unavailable_tcp = in["nat_port_unavailable_tcp"].(int)
		ret.Nat_port_unavailable_udp = in["nat_port_unavailable_udp"].(int)
		ret.Nat_port_unavailable_icmp = in["nat_port_unavailable_icmp"].(int)
		ret.Fullcone_failure = in["fullcone_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobal2576(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobal2576 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobal2576
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsInc2577(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate2578(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsInc2577(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsInc2577 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsInc2577
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NatPortUnavailableTcp = in["nat_port_unavailable_tcp"].(int)
		ret.NatPortUnavailableUdp = in["nat_port_unavailable_udp"].(int)
		ret.NatPortUnavailableIcmp = in["nat_port_unavailable_icmp"].(int)
		ret.SessionUserQuotaExceeded = in["session_user_quota_exceeded"].(int)
		ret.FullconeFailure = in["fullcone_failure"].(int)
		ret.Nat44InboundFiltered = in["nat44_inbound_filtered"].(int)
		ret.Nat64InboundFiltered = in["nat64_inbound_filtered"].(int)
		ret.DsliteInboundFiltered = in["dslite_inbound_filtered"].(int)
		ret.Nat44EifLimitExceeded = in["nat44_eif_limit_exceeded"].(int)
		ret.Nat64EifLimitExceeded = in["nat64_eif_limit_exceeded"].(int)
		ret.DsliteEifLimitExceeded = in["dslite_eif_limit_exceeded"].(int)
		ret.StandbyDrop = in["standby_drop"].(int)
		ret.FixedNatFullconeSelfHairpinningDro = in["fixed_nat_fullcone_self_hairpinning_dro"].(int)
		ret.SixrdDrop = in["sixrd_drop"].(int)
		ret.DestRlistDrop = in["dest_rlist_drop"].(int)
		ret.DestRlistPassThrough = in["dest_rlist_pass_through"].(int)
		ret.DestRlistSnatDrop = in["dest_rlist_snat_drop"].(int)
		ret.ConfigNotFound = in["config_not_found"].(int)
		ret.PortOverloadFailed = in["port_overload_failed"].(int)
		ret.HaSessionUserQuotaExceeded = in["ha_session_user_quota_exceeded"].(int)
		ret.UnusedIpDrop = in["unused_ip_drop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate2578(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate2578 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate2578
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.NatPortUnavailableTcp = in["nat_port_unavailable_tcp"].(int)
		ret.NatPortUnavailableUdp = in["nat_port_unavailable_udp"].(int)
		ret.NatPortUnavailableIcmp = in["nat_port_unavailable_icmp"].(int)
		ret.SessionUserQuotaExceeded = in["session_user_quota_exceeded"].(int)
		ret.FullconeFailure = in["fullcone_failure"].(int)
		ret.Nat44InboundFiltered = in["nat44_inbound_filtered"].(int)
		ret.Nat64InboundFiltered = in["nat64_inbound_filtered"].(int)
		ret.DsliteInboundFiltered = in["dslite_inbound_filtered"].(int)
		ret.Nat44EifLimitExceeded = in["nat44_eif_limit_exceeded"].(int)
		ret.Nat64EifLimitExceeded = in["nat64_eif_limit_exceeded"].(int)
		ret.DsliteEifLimitExceeded = in["dslite_eif_limit_exceeded"].(int)
		ret.StandbyDrop = in["standby_drop"].(int)
		ret.FixedNatFullconeSelfHairpinningDro = in["fixed_nat_fullcone_self_hairpinning_dro"].(int)
		ret.SixrdDrop = in["sixrd_drop"].(int)
		ret.DestRlistDrop = in["dest_rlist_drop"].(int)
		ret.DestRlistPassThrough = in["dest_rlist_pass_through"].(int)
		ret.DestRlistSnatDrop = in["dest_rlist_snat_drop"].(int)
		ret.ConfigNotFound = in["config_not_found"].(int)
		ret.PortOverloadFailed = in["port_overload_failed"].(int)
		ret.HaSessionUserQuotaExceeded = in["ha_session_user_quota_exceeded"].(int)
		ret.UnusedIpDrop = in["unused_ip_drop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgPptp2579(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgPptp2579 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgPptp2579
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsInc2580(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsRate2581(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsInc2580(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsInc2580 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsInc2580
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CallReqPnsCallIdMismatch = in["call_req_pns_call_id_mismatch"].(int)
		ret.CallReplyPnsCallIdMismatch = in["call_reply_pns_call_id_mismatch"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsRate2581(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsRate2581 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsRate2581
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.CallReqPnsCallIdMismatch = in["call_req_pns_call_id_mismatch"].(int)
		ret.CallReplyPnsCallIdMismatch = in["call_reply_pns_call_id_mismatch"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgRtsp2582(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgRtsp2582 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgRtsp2582
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsInc2583(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsRate2584(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsInc2583(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsInc2583 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsInc2583
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StreamCreationFailure = in["stream_creation_failure"].(int)
		ret.PortAllocationFailure = in["port_allocation_failure"].(int)
		ret.NoSessionMem = in["no_session_mem"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsRate2584(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsRate2584 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsRate2584
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.StreamCreationFailure = in["stream_creation_failure"].(int)
		ret.PortAllocationFailure = in["port_allocation_failure"].(int)
		ret.NoSessionMem = in["no_session_mem"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgSip2585(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgSip2585 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgSip2585
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsInc2586(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsRate2587(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsInc2586(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsInc2586 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsInc2586
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MethodUnknown = in["method_unknown"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsRate2587(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsRate2587 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsRate2587
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.MethodUnknown = in["method_unknown"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Pcp2588(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Pcp2588 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Pcp2588
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsInc2589(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsRate2590(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsInc2589(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsInc2589 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsInc2589
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PktNotRequestDrop = in["pkt_not_request_drop"].(int)
		ret.PktTooShortDrop = in["pkt_too_short_drop"].(int)
		ret.NorouteDrop = in["noroute_drop"].(int)
		ret.UnsupportedVersion = in["unsupported_version"].(int)
		ret.NotAuthorized = in["not_authorized"].(int)
		ret.MalformRequest = in["malform_request"].(int)
		ret.UnsuppOpcode = in["unsupp_opcode"].(int)
		ret.UnsuppOption = in["unsupp_option"].(int)
		ret.MalformOption = in["malform_option"].(int)
		ret.NoResources = in["no_resources"].(int)
		ret.UnsuppProtocol = in["unsupp_protocol"].(int)
		ret.CannotProvideSuggest = in["cannot_provide_suggest"].(int)
		ret.AddressMismatch = in["address_mismatch"].(int)
		ret.ExcessiveRemotePeers = in["excessive_remote_peers"].(int)
		ret.PktNotFromNatInside = in["pkt_not_from_nat_inside"].(int)
		ret.L4ProcessError = in["l4_process_error"].(int)
		ret.InternalErrorDrop = in["internal_error_drop"].(int)
		ret.Unsol_ance_sent_fail = in["unsol_ance_sent_fail"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsRate2590(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsRate2590 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsRate2590
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.PktNotRequestDrop = in["pkt_not_request_drop"].(int)
		ret.PktTooShortDrop = in["pkt_too_short_drop"].(int)
		ret.NorouteDrop = in["noroute_drop"].(int)
		ret.UnsupportedVersion = in["unsupported_version"].(int)
		ret.NotAuthorized = in["not_authorized"].(int)
		ret.MalformRequest = in["malform_request"].(int)
		ret.UnsuppOpcode = in["unsupp_opcode"].(int)
		ret.UnsuppOption = in["unsupp_option"].(int)
		ret.MalformOption = in["malform_option"].(int)
		ret.NoResources = in["no_resources"].(int)
		ret.UnsuppProtocol = in["unsupp_protocol"].(int)
		ret.CannotProvideSuggest = in["cannot_provide_suggest"].(int)
		ret.AddressMismatch = in["address_mismatch"].(int)
		ret.ExcessiveRemotePeers = in["excessive_remote_peers"].(int)
		ret.PktNotFromNatInside = in["pkt_not_from_nat_inside"].(int)
		ret.L4ProcessError = in["l4_process_error"].(int)
		ret.InternalErrorDrop = in["internal_error_drop"].(int)
		ret.Unsol_ance_sent_fail = in["unsol_ance_sent_fail"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Logging2591(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Logging2591 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Logging2591
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsInc2592(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsRate2593(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsInc2592(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsInc2592 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsInc2592
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogDropped = in["log_dropped"].(int)
		ret.ConnTcpDropped = in["conn_tcp_dropped"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsRate2593(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsRate2593 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsRate2593
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.LogDropped = in["log_dropped"].(int)
		ret.ConnTcpDropped = in["conn_tcp_dropped"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6L42594(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6L42594 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6L42594
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6L4TriggerStatsInc2595(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6L4TriggerStatsRate2596(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6L4TriggerStatsInc2595(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6L4TriggerStatsInc2595 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6L4TriggerStatsInc2595
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OutOfSessionMemory = in["out_of_session_memory"].(int)
		ret.IcmpHostUnreachableSent = in["icmp_host_unreachable_sent"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6L4TriggerStatsRate2596(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6L4TriggerStatsRate2596 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6L4TriggerStatsRate2596
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.OutOfSessionMemory = in["out_of_session_memory"].(int)
		ret.IcmpHostUnreachableSent = in["icmp_host_unreachable_sent"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Icmp2597(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Icmp2597 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Icmp2597
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsInc2598(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsRate2599(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsInc2598(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsInc2598 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsInc2598
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IcmpToIcmpErr = in["icmp_to_icmp_err"].(int)
		ret.IcmpToIcmpv6Err = in["icmp_to_icmpv6_err"].(int)
		ret.Icmpv6ToIcmpErr = in["icmpv6_to_icmp_err"].(int)
		ret.Icmpv6ToIcmpv6Err = in["icmpv6_to_icmpv6_err"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsRate2599(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsRate2599 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsRate2599
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.IcmpToIcmpErr = in["icmp_to_icmp_err"].(int)
		ret.IcmpToIcmpv6Err = in["icmp_to_icmpv6_err"].(int)
		ret.Icmpv6ToIcmpErr = in["icmpv6_to_icmp_err"].(int)
		ret.Icmpv6ToIcmpv6Err = in["icmpv6_to_icmpv6_err"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlg2600(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlg2600 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlg2600
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsInc2601(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsRate2602(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsInc2601(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsInc2601 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsInc2601
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RadiusRequstDropped = in["radius_requst_dropped"].(int)
		ret.RadiusResponseDropped = in["radius_response_dropped"].(int)
		ret.OutOfMemoryDropped = in["out_of_memory_dropped"].(int)
		ret.QueueLenExceedDropped = in["queue_len_exceed_dropped"].(int)
		ret.OutOfOrderDropped = in["out_of_order_dropped"].(int)
		ret.HeaderInsertionFailed = in["header_insertion_failed"].(int)
		ret.HeaderRemovalFailed = in["header_removal_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsRate2602(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsRate2602 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsRate2602
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.RadiusRequstDropped = in["radius_requst_dropped"].(int)
		ret.RadiusResponseDropped = in["radius_response_dropped"].(int)
		ret.OutOfMemoryDropped = in["out_of_memory_dropped"].(int)
		ret.QueueLenExceedDropped = in["queue_len_exceed_dropped"].(int)
		ret.OutOfOrderDropped = in["out_of_order_dropped"].(int)
		ret.HeaderInsertionFailed = in["header_insertion_failed"].(int)
		ret.HeaderRemovalFailed = in["header_removal_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dns642603(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dns642603 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dns642603
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsInc2604(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsRate2605(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsInc2604(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsInc2604 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsInc2604
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.QueryBadPkt = in["query_bad_pkt"].(int)
		ret.RespBadPkt = in["resp_bad_pkt"].(int)
		ret.RespBadQr = in["resp_bad_qr"].(int)
		ret.Drop = in["drop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsRate2605(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsRate2605 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsRate2605
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.QueryBadPkt = in["query_bad_pkt"].(int)
		ret.RespBadPkt = in["resp_bad_pkt"].(int)
		ret.RespBadQr = in["resp_bad_qr"].(int)
		ret.Drop = in["drop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dhcpv62606(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dhcpv62606 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dhcpv62606
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsInc2607(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsRate2608(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsInc2607(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsInc2607 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsInc2607
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PacketsDropped = in["packets_dropped"].(int)
		ret.PktsDroppedDuringClear = in["pkts_dropped_during_clear"].(int)
		ret.RcvNotSupportedMsg = in["rcv_not_supported_msg"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsRate2608(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsRate2608 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsRate2608
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.PacketsDropped = in["packets_dropped"].(int)
		ret.PktsDroppedDuringClear = in["pkts_dropped_during_clear"].(int)
		ret.RcvNotSupportedMsg = in["rcv_not_supported_msg"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwLogging2609(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwLogging2609 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwLogging2609
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwLoggingTriggerStatsInc2610(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwLoggingTriggerStatsRate2611(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwLoggingTriggerStatsInc2610(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwLoggingTriggerStatsInc2610 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwLoggingTriggerStatsInc2610
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogDropped = in["log_dropped"].(int)
		ret.HttpLoggingInvalidFormat = in["http_logging_invalid_format"].(int)
		ret.SessionLimitExceeded = in["session_limit_exceeded"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwLoggingTriggerStatsRate2611(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwLoggingTriggerStatsRate2611 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwLoggingTriggerStatsRate2611
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.LogDropped = in["log_dropped"].(int)
		ret.HttpLoggingInvalidFormat = in["http_logging_invalid_format"].(int)
		ret.SessionLimitExceeded = in["session_limit_exceeded"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGlobal2612(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGlobal2612 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGlobal2612
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGlobalTriggerStatsInc2613(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGlobalTriggerStatsRate2614(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGlobalTriggerStatsInc2613(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGlobalTriggerStatsInc2613 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGlobalTriggerStatsInc2613
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Fullcone_creation_failure = in["fullcone_creation_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGlobalTriggerStatsRate2614(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGlobalTriggerStatsRate2614 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGlobalTriggerStatsRate2614
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Fullcone_creation_failure = in["fullcone_creation_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgRtsp2615(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgRtsp2615 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgRtsp2615
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgRtspTriggerStatsInc2616(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgRtspTriggerStatsRate2617(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgRtspTriggerStatsInc2616(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgRtspTriggerStatsInc2616 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgRtspTriggerStatsInc2616
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TransportAllocFailure = in["transport_alloc_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgRtspTriggerStatsRate2617(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgRtspTriggerStatsRate2617 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgRtspTriggerStatsRate2617
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.TransportAllocFailure = in["transport_alloc_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptp2618(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptp2618 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptp2618
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsInc2619(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsRate2620(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsInc2619(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsInc2619 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsInc2619
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CallReqPnsCallIdMismatch = in["call_req_pns_call_id_mismatch"].(int)
		ret.CallReplyPnsCallIdMismatch = in["call_reply_pns_call_id_mismatch"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsRate2620(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsRate2620 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsRate2620
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.CallReqPnsCallIdMismatch = in["call_req_pns_call_id_mismatch"].(int)
		ret.CallReplyPnsCallIdMismatch = in["call_reply_pns_call_id_mismatch"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServer2621(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServer2621 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServer2621
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsInc2622(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRate2623(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsInc2622(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsInc2622 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsInc2622
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RadiusRequestDropped = in["radius_request_dropped"].(int)
		ret.RequestBadSecretDropped = in["request_bad_secret_dropped"].(int)
		ret.RequestNoKeyVapDropped = in["request_no_key_vap_dropped"].(int)
		ret.RequestMalformedDropped = in["request_malformed_dropped"].(int)
		ret.RequestIgnored = in["request_ignored"].(int)
		ret.RadiusTableFull = in["radius_table_full"].(int)
		ret.HaStandbyDropped = in["ha_standby_dropped"].(int)
		ret.Ipv6PrefixLengthMismatch = in["ipv6_prefix_length_mismatch"].(int)
		ret.InvalidKey = in["invalid_key"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRate2623(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRate2623 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRate2623
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.RadiusRequestDropped = in["radius_request_dropped"].(int)
		ret.RequestBadSecretDropped = in["request_bad_secret_dropped"].(int)
		ret.RequestNoKeyVapDropped = in["request_no_key_vap_dropped"].(int)
		ret.RequestMalformedDropped = in["request_malformed_dropped"].(int)
		ret.RequestIgnored = in["request_ignored"].(int)
		ret.RadiusTableFull = in["radius_table_full"].(int)
		ret.HaStandbyDropped = in["ha_standby_dropped"].(int)
		ret.Ipv6PrefixLengthMismatch = in["ipv6_prefix_length_mismatch"].(int)
		ret.InvalidKey = in["invalid_key"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookie2624(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookie2624 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookie2624
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsInc2625(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsRate2626(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsInc2625(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsInc2625 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsInc2625
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Verification_failed = in["verification_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsRate2626(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsRate2626 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsRate2626
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Verification_failed = in["verification_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwDdosProtection2627(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwDdosProtection2627 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwDdosProtection2627
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsInc2628(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsRate2629(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsInc2628(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsInc2628 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsInc2628
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entries_too_many = in["ddos_entries_too_many"].(int)
		ret.Ddos_entry_add_to_bgp_failure = in["ddos_entry_add_to_bgp_failure"].(int)
		ret.Ddos_entry_remove_from_bgp_failure = in["ddos_entry_remove_from_bgp_failure"].(int)
		ret.Ddos_packet_dropped = in["ddos_packet_dropped"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsRate2629(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsRate2629 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsRate2629
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Ddos_entries_too_many = in["ddos_entries_too_many"].(int)
		ret.Ddos_entry_add_to_bgp_failure = in["ddos_entry_add_to_bgp_failure"].(int)
		ret.Ddos_entry_remove_from_bgp_failure = in["ddos_entry_remove_from_bgp_failure"].(int)
		ret.Ddos_packet_dropped = in["ddos_packet_dropped"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp2630(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp2630 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp2630
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc2631(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsRate2632(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc2631(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc2631 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc2631
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OutOfSessionMemory = in["out_of_session_memory"].(int)
		ret.BladeOutOfSessionMemory = in["blade_out_of_session_memory"].(int)
		ret.GtpSmpPathCheckFailed = in["gtp_smp_path_check_failed"].(int)
		ret.GtpSmpCCheckFailed = in["gtp_smp_c_check_failed"].(int)
		ret.BladeGtpSmpPathCheckFailed = in["blade_gtp_smp_path_check_failed"].(int)
		ret.BladeGtpSmpCCheckFailed = in["blade_gtp_smp_c_check_failed"].(int)
		ret.GtpTunnelRateLimitEntryCreateFail = in["gtp_tunnel_rate_limit_entry_create_fail"].(int)
		ret.GtpUTunnelRateLimitEntryCreateFa = in["gtp_u_tunnel_rate_limit_entry_create_fa"].(int)
		ret.GtpRateLimitSmpCreateFailure = in["gtp_rate_limit_smp_create_failure"].(int)
		ret.GtpRateLimitT3CtrCreateFailure = in["gtp_rate_limit_t3_ctr_create_failure"].(int)
		ret.GtpRateLimitEntryCreateFailure = in["gtp_rate_limit_entry_create_failure"].(int)
		ret.BladeGtpRateLimitSmpCreateFailure = in["blade_gtp_rate_limit_smp_create_failure"].(int)
		ret.BladeGtpRateLimitT3CtrCreateFail = in["blade_gtp_rate_limit_t3_ctr_create_fail"].(int)
		ret.BladeGtpRateLimitEntryCreateFailu = in["blade_gtp_rate_limit_entry_create_failu"].(int)
		ret.GtpSmpDecSessCountCheckFailed = in["gtp_smp_dec_sess_count_check_failed"].(int)
		ret.GtpUSmpCheckFailed = in["gtp_u_smp_check_failed"].(int)
		ret.GtpInfoExtNotFound = in["gtp_info_ext_not_found"].(int)
		ret.BladeGtpSmpDecSessCountCheckFail = in["blade_gtp_smp_dec_sess_count_check_fail"].(int)
		ret.BladeGtpUSmpCheckFailed = in["blade_gtp_u_smp_check_failed"].(int)
		ret.BladeGtpInfoExtNotFound = in["blade_gtp_info_ext_not_found"].(int)
		ret.BladeGtpSmpSessionCountCheckFaile = in["blade_gtp_smp_session_count_check_faile"].(int)
		ret.GtpCSmpSigCheckFailed = in["gtp_c_smp_sig_check_failed"].(int)
		ret.BladeGtpCSmpSigCheckFailed = in["blade_gtp_c_smp_sig_check_failed"].(int)
		ret.GtpUSmpSigCheckFailed = in["gtp_u_smp_sig_check_failed"].(int)
		ret.BladeGtpUSmpSigCheckFailed = in["blade_gtp_u_smp_sig_check_failed"].(int)
		ret.GtpSmpSigCheckFailed = in["gtp_smp_sig_check_failed"].(int)
		ret.BladeGtpSmpSigCheckFailed = in["blade_gtp_smp_sig_check_failed"].(int)
		ret.GtpCFailConnCreateSlow = in["gtp_c_fail_conn_create_slow"].(int)
		ret.GtpUFailConnCreateSlow = in["gtp_u_fail_conn_create_slow"].(int)
		ret.GtpPathmFailConnCreateSlow = in["gtp_pathm_fail_conn_create_slow"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsRate2632(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsRate2632 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsRate2632
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.OutOfSessionMemory = in["out_of_session_memory"].(int)
		ret.BladeOutOfSessionMemory = in["blade_out_of_session_memory"].(int)
		ret.GtpSmpPathCheckFailed = in["gtp_smp_path_check_failed"].(int)
		ret.GtpSmpCCheckFailed = in["gtp_smp_c_check_failed"].(int)
		ret.BladeGtpSmpPathCheckFailed = in["blade_gtp_smp_path_check_failed"].(int)
		ret.BladeGtpSmpCCheckFailed = in["blade_gtp_smp_c_check_failed"].(int)
		ret.GtpTunnelRateLimitEntryCreateFail = in["gtp_tunnel_rate_limit_entry_create_fail"].(int)
		ret.GtpUTunnelRateLimitEntryCreateFa = in["gtp_u_tunnel_rate_limit_entry_create_fa"].(int)
		ret.GtpRateLimitSmpCreateFailure = in["gtp_rate_limit_smp_create_failure"].(int)
		ret.GtpRateLimitT3CtrCreateFailure = in["gtp_rate_limit_t3_ctr_create_failure"].(int)
		ret.GtpRateLimitEntryCreateFailure = in["gtp_rate_limit_entry_create_failure"].(int)
		ret.BladeGtpRateLimitSmpCreateFailure = in["blade_gtp_rate_limit_smp_create_failure"].(int)
		ret.BladeGtpRateLimitT3CtrCreateFail = in["blade_gtp_rate_limit_t3_ctr_create_fail"].(int)
		ret.BladeGtpRateLimitEntryCreateFailu = in["blade_gtp_rate_limit_entry_create_failu"].(int)
		ret.GtpSmpDecSessCountCheckFailed = in["gtp_smp_dec_sess_count_check_failed"].(int)
		ret.GtpUSmpCheckFailed = in["gtp_u_smp_check_failed"].(int)
		ret.GtpInfoExtNotFound = in["gtp_info_ext_not_found"].(int)
		ret.BladeGtpSmpDecSessCountCheckFail = in["blade_gtp_smp_dec_sess_count_check_fail"].(int)
		ret.BladeGtpUSmpCheckFailed = in["blade_gtp_u_smp_check_failed"].(int)
		ret.BladeGtpInfoExtNotFound = in["blade_gtp_info_ext_not_found"].(int)
		ret.BladeGtpSmpSessionCountCheckFaile = in["blade_gtp_smp_session_count_check_faile"].(int)
		ret.GtpCSmpSigCheckFailed = in["gtp_c_smp_sig_check_failed"].(int)
		ret.BladeGtpCSmpSigCheckFailed = in["blade_gtp_c_smp_sig_check_failed"].(int)
		ret.GtpUSmpSigCheckFailed = in["gtp_u_smp_sig_check_failed"].(int)
		ret.BladeGtpUSmpSigCheckFailed = in["blade_gtp_u_smp_sig_check_failed"].(int)
		ret.GtpSmpSigCheckFailed = in["gtp_smp_sig_check_failed"].(int)
		ret.BladeGtpSmpSigCheckFailed = in["blade_gtp_smp_sig_check_failed"].(int)
		ret.GtpCFailConnCreateSlow = in["gtp_c_fail_conn_create_slow"].(int)
		ret.GtpUFailConnCreateSlow = in["gtp_u_fail_conn_create_slow"].(int)
		ret.GtpPathmFailConnCreateSlow = in["gtp_pathm_fail_conn_create_slow"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemTcp2633(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemTcp2633 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemTcp2633
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemTcpTriggerStatsInc2634(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemTcpTriggerStatsRate2635(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemTcpTriggerStatsInc2634(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemTcpTriggerStatsInc2634 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemTcpTriggerStatsInc2634
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Attemptfails = in["attemptfails"].(int)
		ret.Noroute = in["noroute"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemTcpTriggerStatsRate2635(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemTcpTriggerStatsRate2635 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemTcpTriggerStatsRate2635
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Attemptfails = in["attemptfails"].(int)
		ret.Noroute = in["noroute"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbConnReuse2636(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbConnReuse2636 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbConnReuse2636
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbConnReuseTriggerStatsInc2637(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbConnReuseTriggerStatsRate2638(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbConnReuseTriggerStatsInc2637(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbConnReuseTriggerStatsInc2637 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbConnReuseTriggerStatsInc2637
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ntermi_err = in["ntermi_err"].(int)
		ret.Pause_conn_fail = in["pause_conn_fail"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbConnReuseTriggerStatsRate2638(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbConnReuseTriggerStatsRate2638 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbConnReuseTriggerStatsRate2638
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Ntermi_err = in["ntermi_err"].(int)
		ret.Pause_conn_fail = in["pause_conn_fail"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbAflow2639(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbAflow2639 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbAflow2639
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbAflowTriggerStatsInc2640(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbAflowTriggerStatsRate2641(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbAflowTriggerStatsInc2640(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbAflowTriggerStatsInc2640 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbAflowTriggerStatsInc2640
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Pause_conn_fail = in["pause_conn_fail"].(int)
		ret.Error_resume_conn = in["error_resume_conn"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbAflowTriggerStatsRate2641(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbAflowTriggerStatsRate2641 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbAflowTriggerStatsRate2641
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Pause_conn_fail = in["pause_conn_fail"].(int)
		ret.Error_resume_conn = in["error_resume_conn"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFix2642(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFix2642 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFix2642
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFixTriggerStatsInc2643(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFixTriggerStatsRate2644(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFixTriggerStatsInc2643(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFixTriggerStatsInc2643 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFixTriggerStatsInc2643
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.Noroute = in["noroute"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Client_err = in["client_err"].(int)
		ret.Server_err = in["server_err"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFixTriggerStatsRate2644(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFixTriggerStatsRate2644 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFixTriggerStatsRate2644
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.Noroute = in["noroute"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Client_err = in["client_err"].(int)
		ret.Server_err = in["server_err"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxy2645(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxy2645 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxy2645
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc2646(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsRate2647(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc2646(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc2646 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc2646
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp_err = in["tcp_err"].(int)
		ret.Stream_not_found = in["stream_not_found"].(int)
		ret.Stream_err = in["stream_err"].(int)
		ret.Session_err = in["session_err"].(int)
		ret.Data_no_stream = in["data_no_stream"].(int)
		ret.Data_no_stream_no_goaway = in["data_no_stream_no_goaway"].(int)
		ret.Data_no_stream_goaway_close = in["data_no_stream_goaway_close"].(int)
		ret.Est_cb_no_tuple = in["est_cb_no_tuple"].(int)
		ret.Data_cb_no_tuple = in["data_cb_no_tuple"].(int)
		ret.Ctx_alloc_fail = in["ctx_alloc_fail"].(int)
		ret.Stream_alloc_fail = in["stream_alloc_fail"].(int)
		ret.Http_conn_alloc_fail = in["http_conn_alloc_fail"].(int)
		ret.Request_header_alloc_fail = in["request_header_alloc_fail"].(int)
		ret.Decompress_fail = in["decompress_fail"].(int)
		ret.Invalid_frame_size = in["invalid_frame_size"].(int)
		ret.Invalid_version = in["invalid_version"].(int)
		ret.Compress_ctx_alloc_fail = in["compress_ctx_alloc_fail"].(int)
		ret.Header_compress_fail = in["header_compress_fail"].(int)
		ret.Http_err_stream_closed = in["http_err_stream_closed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsRate2647(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsRate2647 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsRate2647
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Tcp_err = in["tcp_err"].(int)
		ret.Stream_not_found = in["stream_not_found"].(int)
		ret.Stream_err = in["stream_err"].(int)
		ret.Session_err = in["session_err"].(int)
		ret.Data_no_stream = in["data_no_stream"].(int)
		ret.Data_no_stream_no_goaway = in["data_no_stream_no_goaway"].(int)
		ret.Data_no_stream_goaway_close = in["data_no_stream_goaway_close"].(int)
		ret.Est_cb_no_tuple = in["est_cb_no_tuple"].(int)
		ret.Data_cb_no_tuple = in["data_cb_no_tuple"].(int)
		ret.Ctx_alloc_fail = in["ctx_alloc_fail"].(int)
		ret.Stream_alloc_fail = in["stream_alloc_fail"].(int)
		ret.Http_conn_alloc_fail = in["http_conn_alloc_fail"].(int)
		ret.Request_header_alloc_fail = in["request_header_alloc_fail"].(int)
		ret.Decompress_fail = in["decompress_fail"].(int)
		ret.Invalid_frame_size = in["invalid_frame_size"].(int)
		ret.Invalid_version = in["invalid_version"].(int)
		ret.Compress_ctx_alloc_fail = in["compress_ctx_alloc_fail"].(int)
		ret.Header_compress_fail = in["header_compress_fail"].(int)
		ret.Http_err_stream_closed = in["http_err_stream_closed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp22648(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp22648 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp22648
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc2649(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate2650(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc2649(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc2649 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc2649
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Protocol_error = in["protocol_error"].(int)
		ret.Internal_error = in["internal_error"].(int)
		ret.Proxy_alloc_error = in["proxy_alloc_error"].(int)
		ret.Split_buff_fail = in["split_buff_fail"].(int)
		ret.Invalid_frame_size = in["invalid_frame_size"].(int)
		ret.Error_max_invalid_stream = in["error_max_invalid_stream"].(int)
		ret.Data_no_stream = in["data_no_stream"].(int)
		ret.Flow_control_error = in["flow_control_error"].(int)
		ret.Settings_timeout = in["settings_timeout"].(int)
		ret.Frame_size_error = in["frame_size_error"].(int)
		ret.Refused_stream = in["refused_stream"].(int)
		ret.Cancel = in["cancel"].(int)
		ret.Compression_error = in["compression_error"].(int)
		ret.Connect_error = in["connect_error"].(int)
		ret.Enhance_your_calm = in["enhance_your_calm"].(int)
		ret.Inadequate_security = in["inadequate_security"].(int)
		ret.Http_1_1_required = in["http_1_1_required"].(int)
		ret.Deflate_alloc_fail = in["deflate_alloc_fail"].(int)
		ret.Inflate_alloc_fail = in["inflate_alloc_fail"].(int)
		ret.Inflate_header_fail = in["inflate_header_fail"].(int)
		ret.Bad_connection_preface = in["bad_connection_preface"].(int)
		ret.Cant_allocate_control_frame = in["cant_allocate_control_frame"].(int)
		ret.Cant_allocate_settings_frame = in["cant_allocate_settings_frame"].(int)
		ret.Bad_frame_type_for_stream_state = in["bad_frame_type_for_stream_state"].(int)
		ret.Wrong_stream_state = in["wrong_stream_state"].(int)
		ret.Data_queue_alloc_error = in["data_queue_alloc_error"].(int)
		ret.Buff_alloc_error = in["buff_alloc_error"].(int)
		ret.Cant_allocate_rst_frame = in["cant_allocate_rst_frame"].(int)
		ret.Cant_allocate_goaway_frame = in["cant_allocate_goaway_frame"].(int)
		ret.Cant_allocate_ping_frame = in["cant_allocate_ping_frame"].(int)
		ret.Cant_allocate_stream = in["cant_allocate_stream"].(int)
		ret.Cant_allocate_window_frame = in["cant_allocate_window_frame"].(int)
		ret.Header_no_stream = in["header_no_stream"].(int)
		ret.Header_padlen_gt_frame_payload = in["header_padlen_gt_frame_payload"].(int)
		ret.Streams_gt_max_concur_streams = in["streams_gt_max_concur_streams"].(int)
		ret.Idle_state_unexpected_frame = in["idle_state_unexpected_frame"].(int)
		ret.Reserved_local_state_unexpected_frame = in["reserved_local_state_unexpected_frame"].(int)
		ret.Reserved_remote_state_unexpected_frame = in["reserved_remote_state_unexpected_frame"].(int)
		ret.Half_closed_remote_state_unexpected_fra = in["half_closed_remote_state_unexpected_fra"].(int)
		ret.Closed_state_unexpected_frame = in["closed_state_unexpected_frame"].(int)
		ret.Zero_window_size_on_stream = in["zero_window_size_on_stream"].(int)
		ret.Exceeds_max_window_size_stream = in["exceeds_max_window_size_stream"].(int)
		ret.Continuation_before_headers = in["continuation_before_headers"].(int)
		ret.Invalid_frame_during_headers = in["invalid_frame_during_headers"].(int)
		ret.Headers_after_continuation = in["headers_after_continuation"].(int)
		ret.Invalid_push_promise = in["invalid_push_promise"].(int)
		ret.Invalid_stream_id = in["invalid_stream_id"].(int)
		ret.Headers_interleaved = in["headers_interleaved"].(int)
		ret.Trailers_no_end_stream = in["trailers_no_end_stream"].(int)
		ret.Invalid_setting_value = in["invalid_setting_value"].(int)
		ret.Invalid_window_update = in["invalid_window_update"].(int)
		ret.Alloc_fail_total = in["alloc_fail_total"].(int)
		ret.Err_rcvd_total = in["err_rcvd_total"].(int)
		ret.Err_sent_total = in["err_sent_total"].(int)
		ret.Err_sent_proto_err = in["err_sent_proto_err"].(int)
		ret.Err_sent_internal_err = in["err_sent_internal_err"].(int)
		ret.Err_sent_flow_control = in["err_sent_flow_control"].(int)
		ret.Err_sent_setting_timeout = in["err_sent_setting_timeout"].(int)
		ret.Err_sent_stream_closed = in["err_sent_stream_closed"].(int)
		ret.Err_sent_frame_size_err = in["err_sent_frame_size_err"].(int)
		ret.Err_sent_refused_stream = in["err_sent_refused_stream"].(int)
		ret.Err_sent_cancel = in["err_sent_cancel"].(int)
		ret.Err_sent_compression_err = in["err_sent_compression_err"].(int)
		ret.Err_sent_connect_err = in["err_sent_connect_err"].(int)
		ret.Err_sent_your_calm = in["err_sent_your_calm"].(int)
		ret.Err_sent_inadequate_security = in["err_sent_inadequate_security"].(int)
		ret.Err_sent_http11_required = in["err_sent_http11_required"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate2650(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate2650 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate2650
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Protocol_error = in["protocol_error"].(int)
		ret.Internal_error = in["internal_error"].(int)
		ret.Proxy_alloc_error = in["proxy_alloc_error"].(int)
		ret.Split_buff_fail = in["split_buff_fail"].(int)
		ret.Invalid_frame_size = in["invalid_frame_size"].(int)
		ret.Error_max_invalid_stream = in["error_max_invalid_stream"].(int)
		ret.Data_no_stream = in["data_no_stream"].(int)
		ret.Flow_control_error = in["flow_control_error"].(int)
		ret.Settings_timeout = in["settings_timeout"].(int)
		ret.Frame_size_error = in["frame_size_error"].(int)
		ret.Refused_stream = in["refused_stream"].(int)
		ret.Cancel = in["cancel"].(int)
		ret.Compression_error = in["compression_error"].(int)
		ret.Connect_error = in["connect_error"].(int)
		ret.Enhance_your_calm = in["enhance_your_calm"].(int)
		ret.Inadequate_security = in["inadequate_security"].(int)
		ret.Http_1_1_required = in["http_1_1_required"].(int)
		ret.Deflate_alloc_fail = in["deflate_alloc_fail"].(int)
		ret.Inflate_alloc_fail = in["inflate_alloc_fail"].(int)
		ret.Inflate_header_fail = in["inflate_header_fail"].(int)
		ret.Bad_connection_preface = in["bad_connection_preface"].(int)
		ret.Cant_allocate_control_frame = in["cant_allocate_control_frame"].(int)
		ret.Cant_allocate_settings_frame = in["cant_allocate_settings_frame"].(int)
		ret.Bad_frame_type_for_stream_state = in["bad_frame_type_for_stream_state"].(int)
		ret.Wrong_stream_state = in["wrong_stream_state"].(int)
		ret.Data_queue_alloc_error = in["data_queue_alloc_error"].(int)
		ret.Buff_alloc_error = in["buff_alloc_error"].(int)
		ret.Cant_allocate_rst_frame = in["cant_allocate_rst_frame"].(int)
		ret.Cant_allocate_goaway_frame = in["cant_allocate_goaway_frame"].(int)
		ret.Cant_allocate_ping_frame = in["cant_allocate_ping_frame"].(int)
		ret.Cant_allocate_stream = in["cant_allocate_stream"].(int)
		ret.Cant_allocate_window_frame = in["cant_allocate_window_frame"].(int)
		ret.Header_no_stream = in["header_no_stream"].(int)
		ret.Header_padlen_gt_frame_payload = in["header_padlen_gt_frame_payload"].(int)
		ret.Streams_gt_max_concur_streams = in["streams_gt_max_concur_streams"].(int)
		ret.Idle_state_unexpected_frame = in["idle_state_unexpected_frame"].(int)
		ret.Reserved_local_state_unexpected_frame = in["reserved_local_state_unexpected_frame"].(int)
		ret.Reserved_remote_state_unexpected_frame = in["reserved_remote_state_unexpected_frame"].(int)
		ret.Half_closed_remote_state_unexpected_fra = in["half_closed_remote_state_unexpected_fra"].(int)
		ret.Closed_state_unexpected_frame = in["closed_state_unexpected_frame"].(int)
		ret.Zero_window_size_on_stream = in["zero_window_size_on_stream"].(int)
		ret.Exceeds_max_window_size_stream = in["exceeds_max_window_size_stream"].(int)
		ret.Continuation_before_headers = in["continuation_before_headers"].(int)
		ret.Invalid_frame_during_headers = in["invalid_frame_during_headers"].(int)
		ret.Headers_after_continuation = in["headers_after_continuation"].(int)
		ret.Invalid_push_promise = in["invalid_push_promise"].(int)
		ret.Invalid_stream_id = in["invalid_stream_id"].(int)
		ret.Headers_interleaved = in["headers_interleaved"].(int)
		ret.Trailers_no_end_stream = in["trailers_no_end_stream"].(int)
		ret.Invalid_setting_value = in["invalid_setting_value"].(int)
		ret.Invalid_window_update = in["invalid_window_update"].(int)
		ret.Alloc_fail_total = in["alloc_fail_total"].(int)
		ret.Err_rcvd_total = in["err_rcvd_total"].(int)
		ret.Err_sent_total = in["err_sent_total"].(int)
		ret.Err_sent_proto_err = in["err_sent_proto_err"].(int)
		ret.Err_sent_internal_err = in["err_sent_internal_err"].(int)
		ret.Err_sent_flow_control = in["err_sent_flow_control"].(int)
		ret.Err_sent_setting_timeout = in["err_sent_setting_timeout"].(int)
		ret.Err_sent_stream_closed = in["err_sent_stream_closed"].(int)
		ret.Err_sent_frame_size_err = in["err_sent_frame_size_err"].(int)
		ret.Err_sent_refused_stream = in["err_sent_refused_stream"].(int)
		ret.Err_sent_cancel = in["err_sent_cancel"].(int)
		ret.Err_sent_compression_err = in["err_sent_compression_err"].(int)
		ret.Err_sent_connect_err = in["err_sent_connect_err"].(int)
		ret.Err_sent_your_calm = in["err_sent_your_calm"].(int)
		ret.Err_sent_inadequate_security = in["err_sent_inadequate_security"].(int)
		ret.Err_sent_http11_required = in["err_sent_http11_required"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7session2651(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7session2651 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7session2651
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsInc2652(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate2653(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsInc2652(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsInc2652 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsInc2652
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Conn_not_exist = in["conn_not_exist"].(int)
		ret.Wbuf_cb_failed = in["wbuf_cb_failed"].(int)
		ret.Err_event = in["err_event"].(int)
		ret.Err_cb_failed = in["err_cb_failed"].(int)
		ret.Server_conn_failed = in["server_conn_failed"].(int)
		ret.Server_select_fail = in["server_select_fail"].(int)
		ret.Data_cb_failed = in["data_cb_failed"].(int)
		ret.Hps_fwdreq_fail = in["hps_fwdreq_fail"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate2653(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate2653 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate2653
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Conn_not_exist = in["conn_not_exist"].(int)
		ret.Wbuf_cb_failed = in["wbuf_cb_failed"].(int)
		ret.Err_event = in["err_event"].(int)
		ret.Err_cb_failed = in["err_cb_failed"].(int)
		ret.Server_conn_failed = in["server_conn_failed"].(int)
		ret.Server_select_fail = in["server_select_fail"].(int)
		ret.Data_cb_failed = in["data_cb_failed"].(int)
		ret.Hps_fwdreq_fail = in["hps_fwdreq_fail"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmpp2654(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmpp2654 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmpp2654
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmppTriggerStatsInc2655(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmppTriggerStatsRate2656(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmppTriggerStatsInc2655(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmppTriggerStatsInc2655 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmppTriggerStatsInc2655
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Msg_proxy_client_fail = in["msg_proxy_client_fail"].(int)
		ret.Msg_proxy_server_fail = in["msg_proxy_server_fail"].(int)
		ret.Msg_proxy_fail_start_server_conn = in["msg_proxy_fail_start_server_conn"].(int)
		ret.Select_client_fail = in["select_client_fail"].(int)
		ret.Select_server_fail = in["select_server_fail"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmppTriggerStatsRate2656(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmppTriggerStatsRate2656 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmppTriggerStatsRate2656
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Msg_proxy_client_fail = in["msg_proxy_client_fail"].(int)
		ret.Msg_proxy_server_fail = in["msg_proxy_server_fail"].(int)
		ret.Msg_proxy_fail_start_server_conn = in["msg_proxy_fail_start_server_conn"].(int)
		ret.Select_client_fail = in["select_client_fail"].(int)
		ret.Select_server_fail = in["select_server_fail"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmtp2657(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmtp2657 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmtp2657
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmtpTriggerStatsInc2658(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmtpTriggerStatsRate2659(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmtpTriggerStatsInc2658(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmtpTriggerStatsInc2658 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmtpTriggerStatsInc2658
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.No_proxy = in["no_proxy"].(int)
		ret.Parse_req_fail = in["parse_req_fail"].(int)
		ret.Server_select_fail = in["server_select_fail"].(int)
		ret.Forward_req_fail = in["forward_req_fail"].(int)
		ret.Forward_req_data_fail = in["forward_req_data_fail"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Send_client_service_not_ready = in["send_client_service_not_ready"].(int)
		ret.Recv_server_unknow_reply_code = in["recv_server_unknow_reply_code"].(int)
		ret.Read_request_line_fail = in["read_request_line_fail"].(int)
		ret.Get_all_headers_fail = in["get_all_headers_fail"].(int)
		ret.Too_many_headers = in["too_many_headers"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Line_extend_fail = in["line_extend_fail"].(int)
		ret.Line_table_extend_fail = in["line_table_extend_fail"].(int)
		ret.Parse_request_line_fail = in["parse_request_line_fail"].(int)
		ret.Insert_resonse_line_fail = in["insert_resonse_line_fail"].(int)
		ret.Remove_resonse_line_fail = in["remove_resonse_line_fail"].(int)
		ret.Parse_resonse_line_fail = in["parse_resonse_line_fail"].(int)
		ret.Server_starttls_fail = in["server_starttls_fail"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmtpTriggerStatsRate2659(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmtpTriggerStatsRate2659 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmtpTriggerStatsRate2659
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.No_proxy = in["no_proxy"].(int)
		ret.Parse_req_fail = in["parse_req_fail"].(int)
		ret.Server_select_fail = in["server_select_fail"].(int)
		ret.Forward_req_fail = in["forward_req_fail"].(int)
		ret.Forward_req_data_fail = in["forward_req_data_fail"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Send_client_service_not_ready = in["send_client_service_not_ready"].(int)
		ret.Recv_server_unknow_reply_code = in["recv_server_unknow_reply_code"].(int)
		ret.Read_request_line_fail = in["read_request_line_fail"].(int)
		ret.Get_all_headers_fail = in["get_all_headers_fail"].(int)
		ret.Too_many_headers = in["too_many_headers"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Line_extend_fail = in["line_extend_fail"].(int)
		ret.Line_table_extend_fail = in["line_table_extend_fail"].(int)
		ret.Parse_request_line_fail = in["parse_request_line_fail"].(int)
		ret.Insert_resonse_line_fail = in["insert_resonse_line_fail"].(int)
		ret.Remove_resonse_line_fail = in["remove_resonse_line_fail"].(int)
		ret.Parse_resonse_line_fail = in["parse_resonse_line_fail"].(int)
		ret.Server_starttls_fail = in["server_starttls_fail"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMqtt2660(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMqtt2660 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMqtt2660
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMqttTriggerStatsInc2661(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMqttTriggerStatsRate2662(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMqttTriggerStatsInc2661(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMqttTriggerStatsInc2661 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMqttTriggerStatsInc2661
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Parse_connect_fail = in["parse_connect_fail"].(int)
		ret.Parse_publish_fail = in["parse_publish_fail"].(int)
		ret.Parse_subscribe_fail = in["parse_subscribe_fail"].(int)
		ret.Parse_unsubscribe_fail = in["parse_unsubscribe_fail"].(int)
		ret.Tuple_not_linked = in["tuple_not_linked"].(int)
		ret.Tuple_already_linked = in["tuple_already_linked"].(int)
		ret.Conn_null = in["conn_null"].(int)
		ret.Client_id_null = in["client_id_null"].(int)
		ret.Session_exist = in["session_exist"].(int)
		ret.Insertion_failed = in["insertion_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMqttTriggerStatsRate2662(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMqttTriggerStatsRate2662 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMqttTriggerStatsRate2662
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Parse_connect_fail = in["parse_connect_fail"].(int)
		ret.Parse_publish_fail = in["parse_publish_fail"].(int)
		ret.Parse_subscribe_fail = in["parse_subscribe_fail"].(int)
		ret.Parse_unsubscribe_fail = in["parse_unsubscribe_fail"].(int)
		ret.Tuple_not_linked = in["tuple_not_linked"].(int)
		ret.Tuple_already_linked = in["tuple_already_linked"].(int)
		ret.Conn_null = in["conn_null"].(int)
		ret.Client_id_null = in["client_id_null"].(int)
		ret.Session_exist = in["session_exist"].(int)
		ret.Insertion_failed = in["insertion_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcap2663(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcap2663 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcap2663
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsInc2664(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRate2665(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsInc2664(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsInc2664 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsInc2664
	if count1 > 0 {
		in := d[0].(map[string]interface{})
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
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRate2665(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRate2665 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRate2665
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
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
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSip2666(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSip2666 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSip2666
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSipTriggerStatsInc2667(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSipTriggerStatsRate2668(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSipTriggerStatsInc2667(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSipTriggerStatsInc2667 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSipTriggerStatsInc2667
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Msg_proxy_client_fail = in["msg_proxy_client_fail"].(int)
		ret.Msg_proxy_server_fail = in["msg_proxy_server_fail"].(int)
		ret.Msg_proxy_fail_start_server_conn = in["msg_proxy_fail_start_server_conn"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSipTriggerStatsRate2668(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSipTriggerStatsRate2668 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSipTriggerStatsRate2668
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Msg_proxy_client_fail = in["msg_proxy_client_fail"].(int)
		ret.Msg_proxy_server_fail = in["msg_proxy_server_fail"].(int)
		ret.Msg_proxy_fail_start_server_conn = in["msg_proxy_fail_start_server_conn"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHwCompress2669(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHwCompress2669 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHwCompress2669
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHwCompressTriggerStatsInc2670(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHwCompressTriggerStatsRate2671(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHwCompressTriggerStatsInc2670(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHwCompressTriggerStatsInc2670 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHwCompressTriggerStatsInc2670
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Failure_count = in["failure_count"].(int)
		ret.Failure_code = in["failure_code"].(int)
		ret.Ring_full_count = in["ring_full_count"].(int)
		ret.Max_outstanding_request_count = in["max_outstanding_request_count"].(int)
		ret.Max_outstanding_submit_count = in["max_outstanding_submit_count"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHwCompressTriggerStatsRate2671(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHwCompressTriggerStatsRate2671 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHwCompressTriggerStatsRate2671
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Failure_count = in["failure_count"].(int)
		ret.Failure_code = in["failure_code"].(int)
		ret.Ring_full_count = in["ring_full_count"].(int)
		ret.Max_outstanding_request_count = in["max_outstanding_request_count"].(int)
		ret.Max_outstanding_submit_count = in["max_outstanding_submit_count"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMysql2672(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMysql2672 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMysql2672
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMysqlTriggerStatsInc2673(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMysqlTriggerStatsRate2674(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMysqlTriggerStatsInc2673(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMysqlTriggerStatsInc2673 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMysqlTriggerStatsInc2673
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Session_err = in["session_err"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMysqlTriggerStatsRate2674(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMysqlTriggerStatsRate2674 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMysqlTriggerStatsRate2674
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Session_err = in["session_err"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMssql2675(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMssql2675 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMssql2675
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMssqlTriggerStatsInc2676(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMssqlTriggerStatsRate2677(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMssqlTriggerStatsInc2676(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMssqlTriggerStatsInc2676 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMssqlTriggerStatsInc2676
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Session_err = in["session_err"].(int)
		ret.Auth_failure = in["auth_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMssqlTriggerStatsRate2677(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMssqlTriggerStatsRate2677 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMssqlTriggerStatsRate2677
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Session_err = in["session_err"].(int)
		ret.Auth_failure = in["auth_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbCrlSrcip2678(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbCrlSrcip2678 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbCrlSrcip2678
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsInc2679(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsRate2680(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsInc2679(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsInc2679 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsInc2679
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Out_of_sessions = in["out_of_sessions"].(int)
		ret.Too_many_sessions = in["too_many_sessions"].(int)
		ret.Threshold_exceed = in["threshold_exceed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsRate2680(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsRate2680 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsRate2680
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Out_of_sessions = in["out_of_sessions"].(int)
		ret.Too_many_sessions = in["too_many_sessions"].(int)
		ret.Threshold_exceed = in["threshold_exceed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGeneric2681(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGeneric2681 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGeneric2681
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGenericTriggerStatsInc2682(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGenericTriggerStatsRate2683(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGenericTriggerStatsInc2682(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGenericTriggerStatsInc2682 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGenericTriggerStatsInc2682
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Client_fail = in["client_fail"].(int)
		ret.Server_fail = in["server_fail"].(int)
		ret.Mismatch_fwd_id = in["mismatch_fwd_id"].(int)
		ret.Mismatch_rev_id = in["mismatch_rev_id"].(int)
		ret.Unkwn_cmd_code = in["unkwn_cmd_code"].(int)
		ret.No_session_id = in["no_session_id"].(int)
		ret.No_fwd_tuple = in["no_fwd_tuple"].(int)
		ret.No_rev_tuple = in["no_rev_tuple"].(int)
		ret.Dcmsg_error = in["dcmsg_error"].(int)
		ret.Retry_client_request_fail = in["retry_client_request_fail"].(int)
		ret.Reply_unknown_session_id = in["reply_unknown_session_id"].(int)
		ret.Client_select_fail = in["client_select_fail"].(int)
		ret.Invalid_avp = in["invalid_avp"].(int)
		ret.Reply_error_info_fail = in["reply_error_info_fail"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGenericTriggerStatsRate2683(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGenericTriggerStatsRate2683 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGenericTriggerStatsRate2683
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Client_fail = in["client_fail"].(int)
		ret.Server_fail = in["server_fail"].(int)
		ret.Mismatch_fwd_id = in["mismatch_fwd_id"].(int)
		ret.Mismatch_rev_id = in["mismatch_rev_id"].(int)
		ret.Unkwn_cmd_code = in["unkwn_cmd_code"].(int)
		ret.No_session_id = in["no_session_id"].(int)
		ret.No_fwd_tuple = in["no_fwd_tuple"].(int)
		ret.No_rev_tuple = in["no_rev_tuple"].(int)
		ret.Dcmsg_error = in["dcmsg_error"].(int)
		ret.Retry_client_request_fail = in["retry_client_request_fail"].(int)
		ret.Reply_unknown_session_id = in["reply_unknown_session_id"].(int)
		ret.Client_select_fail = in["client_select_fail"].(int)
		ret.Invalid_avp = in["invalid_avp"].(int)
		ret.Reply_error_info_fail = in["reply_error_info_fail"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist2684(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist2684 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist2684
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsInc2685(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate2686(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsInc2685(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsInc2685 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsInc2685
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hash_tbl_trylock_fail = in["hash_tbl_trylock_fail"].(int)
		ret.Hash_tbl_create_fail = in["hash_tbl_create_fail"].(int)
		ret.Hash_tbl_rst_updown = in["hash_tbl_rst_updown"].(int)
		ret.Hash_tbl_rst_adddel = in["hash_tbl_rst_adddel"].(int)
		ret.Url_hash_fail = in["url_hash_fail"].(int)
		ret.Header_hash_fail = in["header_hash_fail"].(int)
		ret.Src_ip_fail = in["src_ip_fail"].(int)
		ret.Src_ip_new_sess_cache_fail = in["src_ip_new_sess_cache_fail"].(int)
		ret.Src_ip_new_sess_sel_fail = in["src_ip_new_sess_sel_fail"].(int)
		ret.Src_ip_hash_fail = in["src_ip_hash_fail"].(int)
		ret.Dst_ip_fail = in["dst_ip_fail"].(int)
		ret.Dst_ip_new_sess_cache_fail = in["dst_ip_new_sess_cache_fail"].(int)
		ret.Dst_ip_new_sess_sel_fail = in["dst_ip_new_sess_sel_fail"].(int)
		ret.Dst_ip_hash_fail = in["dst_ip_hash_fail"].(int)
		ret.Cssl_sid_not_found = in["cssl_sid_not_found"].(int)
		ret.Cssl_sid_not_match = in["cssl_sid_not_match"].(int)
		ret.Sssl_sid_not_found = in["sssl_sid_not_found"].(int)
		ret.Sssl_sid_not_match = in["sssl_sid_not_match"].(int)
		ret.Ssl_sid_persist_fail = in["ssl_sid_persist_fail"].(int)
		ret.Ssl_sid_session_fail = in["ssl_sid_session_fail"].(int)
		ret.Cookie_persist_fail = in["cookie_persist_fail"].(int)
		ret.Cookie_not_found = in["cookie_not_found"].(int)
		ret.Cookie_invalid = in["cookie_invalid"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate2686(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate2686 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate2686
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Hash_tbl_trylock_fail = in["hash_tbl_trylock_fail"].(int)
		ret.Hash_tbl_create_fail = in["hash_tbl_create_fail"].(int)
		ret.Hash_tbl_rst_updown = in["hash_tbl_rst_updown"].(int)
		ret.Hash_tbl_rst_adddel = in["hash_tbl_rst_adddel"].(int)
		ret.Url_hash_fail = in["url_hash_fail"].(int)
		ret.Header_hash_fail = in["header_hash_fail"].(int)
		ret.Src_ip_fail = in["src_ip_fail"].(int)
		ret.Src_ip_new_sess_cache_fail = in["src_ip_new_sess_cache_fail"].(int)
		ret.Src_ip_new_sess_sel_fail = in["src_ip_new_sess_sel_fail"].(int)
		ret.Src_ip_hash_fail = in["src_ip_hash_fail"].(int)
		ret.Dst_ip_fail = in["dst_ip_fail"].(int)
		ret.Dst_ip_new_sess_cache_fail = in["dst_ip_new_sess_cache_fail"].(int)
		ret.Dst_ip_new_sess_sel_fail = in["dst_ip_new_sess_sel_fail"].(int)
		ret.Dst_ip_hash_fail = in["dst_ip_hash_fail"].(int)
		ret.Cssl_sid_not_found = in["cssl_sid_not_found"].(int)
		ret.Cssl_sid_not_match = in["cssl_sid_not_match"].(int)
		ret.Sssl_sid_not_found = in["sssl_sid_not_found"].(int)
		ret.Sssl_sid_not_match = in["sssl_sid_not_match"].(int)
		ret.Ssl_sid_persist_fail = in["ssl_sid_persist_fail"].(int)
		ret.Ssl_sid_session_fail = in["ssl_sid_session_fail"].(int)
		ret.Cookie_persist_fail = in["cookie_persist_fail"].(int)
		ret.Cookie_not_found = in["cookie_not_found"].(int)
		ret.Cookie_invalid = in["cookie_invalid"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttpProxy2687(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttpProxy2687 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttpProxy2687
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsInc2688(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsRate2689(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsInc2688(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsInc2688 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsInc2688
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Parsereq_fail = in["parsereq_fail"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.Fwdreq_fail = in["fwdreq_fail"].(int)
		ret.Fwdreqdata_fail = in["fwdreqdata_fail"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Req_over_limit = in["req_over_limit"].(int)
		ret.Req_rate_over_limit = in["req_rate_over_limit"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsRate2689(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsRate2689 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsRate2689
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Parsereq_fail = in["parsereq_fail"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.Fwdreq_fail = in["fwdreq_fail"].(int)
		ret.Fwdreqdata_fail = in["fwdreqdata_fail"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Req_over_limit = in["req_over_limit"].(int)
		ret.Req_rate_over_limit = in["req_rate_over_limit"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL42690(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL42690 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL42690
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsInc2691(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsRate2692(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsInc2691(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsInc2691 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsInc2691
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Syncookiessentfailed = in["syncookiessentfailed"].(int)
		ret.Svrselfail = in["svrselfail"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Snat_no_fwd_route = in["snat_no_fwd_route"].(int)
		ret.Snat_no_rev_route = in["snat_no_rev_route"].(int)
		ret.Snat_icmp_error_process = in["snat_icmp_error_process"].(int)
		ret.Snat_icmp_no_match = in["snat_icmp_no_match"].(int)
		ret.Smart_nat_id_mismatch = in["smart_nat_id_mismatch"].(int)
		ret.Syncookiescheckfailed = in["syncookiescheckfailed"].(int)
		ret.Connlimit_drop = in["connlimit_drop"].(int)
		ret.Conn_rate_limit_drop = in["conn_rate_limit_drop"].(int)
		ret.Conn_rate_limit_reset = in["conn_rate_limit_reset"].(int)
		ret.Dns_policy_drop = in["dns_policy_drop"].(int)
		ret.No_resourse_drop = in["no_resourse_drop"].(int)
		ret.Bw_rate_limit_exceed = in["bw_rate_limit_exceed"].(int)
		ret.L4_cps_exceed = in["l4_cps_exceed"].(int)
		ret.Nat_cps_exceed = in["nat_cps_exceed"].(int)
		ret.L7_cps_exceed = in["l7_cps_exceed"].(int)
		ret.Ssl_cps_exceed = in["ssl_cps_exceed"].(int)
		ret.Ssl_tpt_exceed = in["ssl_tpt_exceed"].(int)
		ret.Concurrent_conn_exceed = in["concurrent_conn_exceed"].(int)
		ret.Svr_syn_handshake_fail = in["svr_syn_handshake_fail"].(int)
		ret.Synattack = in["synattack"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsRate2692(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsRate2692 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsRate2692
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Syncookiessentfailed = in["syncookiessentfailed"].(int)
		ret.Svrselfail = in["svrselfail"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Snat_no_fwd_route = in["snat_no_fwd_route"].(int)
		ret.Snat_no_rev_route = in["snat_no_rev_route"].(int)
		ret.Snat_icmp_error_process = in["snat_icmp_error_process"].(int)
		ret.Snat_icmp_no_match = in["snat_icmp_no_match"].(int)
		ret.Smart_nat_id_mismatch = in["smart_nat_id_mismatch"].(int)
		ret.Syncookiescheckfailed = in["syncookiescheckfailed"].(int)
		ret.Connlimit_drop = in["connlimit_drop"].(int)
		ret.Conn_rate_limit_drop = in["conn_rate_limit_drop"].(int)
		ret.Conn_rate_limit_reset = in["conn_rate_limit_reset"].(int)
		ret.Dns_policy_drop = in["dns_policy_drop"].(int)
		ret.No_resourse_drop = in["no_resourse_drop"].(int)
		ret.Bw_rate_limit_exceed = in["bw_rate_limit_exceed"].(int)
		ret.L4_cps_exceed = in["l4_cps_exceed"].(int)
		ret.Nat_cps_exceed = in["nat_cps_exceed"].(int)
		ret.L7_cps_exceed = in["l7_cps_exceed"].(int)
		ret.Ssl_cps_exceed = in["ssl_cps_exceed"].(int)
		ret.Ssl_tpt_exceed = in["ssl_tpt_exceed"].(int)
		ret.Concurrent_conn_exceed = in["concurrent_conn_exceed"].(int)
		ret.Svr_syn_handshake_fail = in["svr_syn_handshake_fail"].(int)
		ret.Synattack = in["synattack"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFastHttp2693(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFastHttp2693 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFastHttp2693
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFastHttpTriggerStatsInc2694(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFastHttpTriggerStatsRate2695(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFastHttpTriggerStatsInc2694(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFastHttpTriggerStatsInc2694 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFastHttpTriggerStatsInc2694
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Parsereq_fail = in["parsereq_fail"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.Fwdreq_fail = in["fwdreq_fail"].(int)
		ret.Fwdreqdata_fail = in["fwdreqdata_fail"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Full_proxy_fpga_err = in["full_proxy_fpga_err"].(int)
		ret.Req_over_limit = in["req_over_limit"].(int)
		ret.Req_rate_over_limit = in["req_rate_over_limit"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFastHttpTriggerStatsRate2695(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFastHttpTriggerStatsRate2695 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFastHttpTriggerStatsRate2695
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Parsereq_fail = in["parsereq_fail"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.Fwdreq_fail = in["fwdreq_fail"].(int)
		ret.Fwdreqdata_fail = in["fwdreqdata_fail"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Full_proxy_fpga_err = in["full_proxy_fpga_err"].(int)
		ret.Req_over_limit = in["req_over_limit"].(int)
		ret.Req_rate_over_limit = in["req_rate_over_limit"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFtpProxy2696(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFtpProxy2696 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFtpProxy2696
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsInc2697(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsRate2698(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsInc2697(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsInc2697 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsInc2697
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Invalid_start_line = in["invalid_start_line"].(int)
		ret.Smp_create_fail = in["smp_create_fail"].(int)
		ret.Data_server_conn_fail = in["data_server_conn_fail"].(int)
		ret.Data_send_fail = in["data_send_fail"].(int)
		ret.Unsupported_pbsz_value = in["unsupported_pbsz_value"].(int)
		ret.Unsupported_prot_value = in["unsupported_prot_value"].(int)
		ret.Unsupported_command = in["unsupported_command"].(int)
		ret.Bad_sequence = in["bad_sequence"].(int)
		ret.Rsv_persist_conn_fail = in["rsv_persist_conn_fail"].(int)
		ret.Smp_v6_fail = in["smp_v6_fail"].(int)
		ret.Smp_v4_fail = in["smp_v4_fail"].(int)
		ret.Insert_tuple_fail = in["insert_tuple_fail"].(int)
		ret.Cl_est_err = in["cl_est_err"].(int)
		ret.Ser_connecting_err = in["ser_connecting_err"].(int)
		ret.Server_response_err = in["server_response_err"].(int)
		ret.Cl_request_err = in["cl_request_err"].(int)
		ret.Data_conn_start_err = in["data_conn_start_err"].(int)
		ret.Data_serv_connecting_err = in["data_serv_connecting_err"].(int)
		ret.Data_serv_connected_err = in["data_serv_connected_err"].(int)
		ret.Auth_fail = in["auth_fail"].(int)
		ret.Ds_fail = in["ds_fail"].(int)
		ret.Cant_find_port = in["cant_find_port"].(int)
		ret.Cant_find_eprt = in["cant_find_eprt"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsRate2698(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsRate2698 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsRate2698
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Invalid_start_line = in["invalid_start_line"].(int)
		ret.Smp_create_fail = in["smp_create_fail"].(int)
		ret.Data_server_conn_fail = in["data_server_conn_fail"].(int)
		ret.Data_send_fail = in["data_send_fail"].(int)
		ret.Unsupported_pbsz_value = in["unsupported_pbsz_value"].(int)
		ret.Unsupported_prot_value = in["unsupported_prot_value"].(int)
		ret.Unsupported_command = in["unsupported_command"].(int)
		ret.Bad_sequence = in["bad_sequence"].(int)
		ret.Rsv_persist_conn_fail = in["rsv_persist_conn_fail"].(int)
		ret.Smp_v6_fail = in["smp_v6_fail"].(int)
		ret.Smp_v4_fail = in["smp_v4_fail"].(int)
		ret.Insert_tuple_fail = in["insert_tuple_fail"].(int)
		ret.Cl_est_err = in["cl_est_err"].(int)
		ret.Ser_connecting_err = in["ser_connecting_err"].(int)
		ret.Server_response_err = in["server_response_err"].(int)
		ret.Cl_request_err = in["cl_request_err"].(int)
		ret.Data_conn_start_err = in["data_conn_start_err"].(int)
		ret.Data_serv_connecting_err = in["data_serv_connecting_err"].(int)
		ret.Data_serv_connected_err = in["data_serv_connected_err"].(int)
		ret.Auth_fail = in["auth_fail"].(int)
		ret.Ds_fail = in["ds_fail"].(int)
		ret.Cant_find_port = in["cant_find_port"].(int)
		ret.Cant_find_eprt = in["cant_find_eprt"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy2699(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy2699 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy2699
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsInc2700(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsRate2701(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsInc2700(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsInc2700 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsInc2700
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Invalid_start_line = in["invalid_start_line"].(int)
		ret.Cant_find_pasv = in["cant_find_pasv"].(int)
		ret.Smp_create_fail = in["smp_create_fail"].(int)
		ret.Data_server_conn_fail = in["data_server_conn_fail"].(int)
		ret.Data_send_fail = in["data_send_fail"].(int)
		ret.Cant_find_epsv = in["cant_find_epsv"].(int)
		ret.Auth_unsupported = in["auth_unsupported"].(int)
		ret.Unsupported_pbsz_value = in["unsupported_pbsz_value"].(int)
		ret.Unsupported_prot_value = in["unsupported_prot_value"].(int)
		ret.Bad_sequence = in["bad_sequence"].(int)
		ret.Rsv_persist_conn_fail = in["rsv_persist_conn_fail"].(int)
		ret.Smp_v6_fail = in["smp_v6_fail"].(int)
		ret.Smp_v4_fail = in["smp_v4_fail"].(int)
		ret.Insert_tuple_fail = in["insert_tuple_fail"].(int)
		ret.Cl_est_err = in["cl_est_err"].(int)
		ret.Ser_connecting_err = in["ser_connecting_err"].(int)
		ret.Server_response_err = in["server_response_err"].(int)
		ret.Cl_request_err = in["cl_request_err"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsRate2701(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsRate2701 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsRate2701
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Invalid_start_line = in["invalid_start_line"].(int)
		ret.Cant_find_pasv = in["cant_find_pasv"].(int)
		ret.Smp_create_fail = in["smp_create_fail"].(int)
		ret.Data_server_conn_fail = in["data_server_conn_fail"].(int)
		ret.Data_send_fail = in["data_send_fail"].(int)
		ret.Cant_find_epsv = in["cant_find_epsv"].(int)
		ret.Auth_unsupported = in["auth_unsupported"].(int)
		ret.Unsupported_pbsz_value = in["unsupported_pbsz_value"].(int)
		ret.Unsupported_prot_value = in["unsupported_prot_value"].(int)
		ret.Bad_sequence = in["bad_sequence"].(int)
		ret.Rsv_persist_conn_fail = in["rsv_persist_conn_fail"].(int)
		ret.Smp_v6_fail = in["smp_v6_fail"].(int)
		ret.Smp_v4_fail = in["smp_v4_fail"].(int)
		ret.Insert_tuple_fail = in["insert_tuple_fail"].(int)
		ret.Cl_est_err = in["cl_est_err"].(int)
		ret.Ser_connecting_err = in["ser_connecting_err"].(int)
		ret.Server_response_err = in["server_response_err"].(int)
		ret.Cl_request_err = in["cl_request_err"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPop3Proxy2702(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPop3Proxy2702 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPop3Proxy2702
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsInc2703(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsRate2704(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsInc2703(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsInc2703 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsInc2703
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Invalid_start_line = in["invalid_start_line"].(int)
		ret.Unsupported_command = in["unsupported_command"].(int)
		ret.Bad_sequence = in["bad_sequence"].(int)
		ret.Rsv_persist_conn_fail = in["rsv_persist_conn_fail"].(int)
		ret.Smp_v6_fail = in["smp_v6_fail"].(int)
		ret.Smp_v4_fail = in["smp_v4_fail"].(int)
		ret.Insert_tuple_fail = in["insert_tuple_fail"].(int)
		ret.Cl_est_err = in["cl_est_err"].(int)
		ret.Ser_connecting_err = in["ser_connecting_err"].(int)
		ret.Server_response_err = in["server_response_err"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsRate2704(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsRate2704 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsRate2704
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Invalid_start_line = in["invalid_start_line"].(int)
		ret.Unsupported_command = in["unsupported_command"].(int)
		ret.Bad_sequence = in["bad_sequence"].(int)
		ret.Rsv_persist_conn_fail = in["rsv_persist_conn_fail"].(int)
		ret.Smp_v6_fail = in["smp_v6_fail"].(int)
		ret.Smp_v4_fail = in["smp_v4_fail"].(int)
		ret.Insert_tuple_fail = in["insert_tuple_fail"].(int)
		ret.Cl_est_err = in["cl_est_err"].(int)
		ret.Ser_connecting_err = in["ser_connecting_err"].(int)
		ret.Server_response_err = in["server_response_err"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSwitch2705(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSwitch2705 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSwitch2705
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSwitchTriggerStatsInc2706(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSwitchTriggerStatsRate2707(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSwitchTriggerStatsInc2706(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSwitchTriggerStatsInc2706 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSwitchTriggerStatsInc2706
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Lacp_tx_intf_err_drop = in["lacp_tx_intf_err_drop"].(int)
		ret.Unnumbered_nat_error = in["unnumbered_nat_error"].(int)
		ret.Unnumbered_unsupported_drop = in["unnumbered_unsupported_drop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSwitchTriggerStatsRate2707(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSwitchTriggerStatsRate2707 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSwitchTriggerStatsRate2707
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Lacp_tx_intf_err_drop = in["lacp_tx_intf_err_drop"].(int)
		ret.Unnumbered_nat_error = in["unnumbered_nat_error"].(int)
		ret.Unnumbered_unsupported_drop = in["unnumbered_unsupported_drop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRcCache2708(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRcCache2708 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRcCache2708
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRcCacheTriggerStatsInc2709(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRcCacheTriggerStatsRate2710(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRcCacheTriggerStatsInc2709(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRcCacheTriggerStatsInc2709 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRcCacheTriggerStatsInc2709
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rv_failure = in["rv_failure"].(int)
		ret.Content_toobig = in["content_toobig"].(int)
		ret.Content_toosmall = in["content_toosmall"].(int)
		ret.Entry_create_failures = in["entry_create_failures"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRcCacheTriggerStatsRate2710(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRcCacheTriggerStatsRate2710 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRcCacheTriggerStatsRate2710
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Rv_failure = in["rv_failure"].(int)
		ret.Content_toobig = in["content_toobig"].(int)
		ret.Content_toosmall = in["content_toosmall"].(int)
		ret.Entry_create_failures = in["entry_create_failures"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCounters2711(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCounters2711 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCounters2711
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsInc2712(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsRate2713(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsInc2712(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsInc2712 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsInc2712
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.So_pkts_slb_nat_reserve_fail = in["so_pkts_slb_nat_reserve_fail"].(int)
		ret.So_pkts_slb_nat_release_fail = in["so_pkts_slb_nat_release_fail"].(int)
		ret.So_pkts_l2redirect_dest_mac_zero_drop = in["so_pkts_l2redirect_dest_mac_zero_drop"].(int)
		ret.So_pkts_l2redirect_interface_not_up = in["so_pkts_l2redirect_interface_not_up"].(int)
		ret.So_pkts_l2redirect_invalid_redirect_inf = in["so_pkts_l2redirect_invalid_redirect_inf"].(int)
		ret.So_pkts_l3_redirect_encap_error_drop = in["so_pkts_l3_redirect_encap_error_drop"].(int)
		ret.So_pkts_l3_redirect_inner_mac_zero_drop = in["so_pkts_l3_redirect_inner_mac_zero_drop"].(int)
		ret.So_pkts_l3_redirect_table_error = in["so_pkts_l3_redirect_table_error"].(int)
		ret.So_pkts_l3_redirect_fragmentation_error = in["so_pkts_l3_redirect_fragmentation_error"].(int)
		ret.So_pkts_l3_redirect_table_no_entry_foun = in["so_pkts_l3_redirect_table_no_entry_foun"].(int)
		ret.So_pkts_l3_redirect_invalid_dev_dir = in["so_pkts_l3_redirect_invalid_dev_dir"].(int)
		ret.So_pkts_l3_redirect_chassis_dest_mac_er = in["so_pkts_l3_redirect_chassis_dest_mac_er"].(int)
		ret.So_pkts_l2redirect_vlan_retrieval_error = in["so_pkts_l2redirect_vlan_retrieval_error"].(int)
		ret.So_pkts_l2redirect_port_retrieval_error = in["so_pkts_l2redirect_port_retrieval_error"].(int)
		ret.So_pkts_l2redirect_loop_detect_drop = in["so_pkts_l2redirect_loop_detect_drop"].(int)
		ret.So_pkts_l2redirect_frag_vlan_retrieval_ = in["so_pkts_l2redirect_frag_vlan_retrieval_"].(int)
		ret.So_pkts_l2redirect_tx_frag_vlan_add_fai = in["so_pkts_l2redirect_tx_frag_vlan_add_fai"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsRate2713(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsRate2713 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsRate2713
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.So_pkts_slb_nat_reserve_fail = in["so_pkts_slb_nat_reserve_fail"].(int)
		ret.So_pkts_slb_nat_release_fail = in["so_pkts_slb_nat_release_fail"].(int)
		ret.So_pkts_l2redirect_dest_mac_zero_drop = in["so_pkts_l2redirect_dest_mac_zero_drop"].(int)
		ret.So_pkts_l2redirect_interface_not_up = in["so_pkts_l2redirect_interface_not_up"].(int)
		ret.So_pkts_l2redirect_invalid_redirect_inf = in["so_pkts_l2redirect_invalid_redirect_inf"].(int)
		ret.So_pkts_l3_redirect_encap_error_drop = in["so_pkts_l3_redirect_encap_error_drop"].(int)
		ret.So_pkts_l3_redirect_inner_mac_zero_drop = in["so_pkts_l3_redirect_inner_mac_zero_drop"].(int)
		ret.So_pkts_l3_redirect_table_error = in["so_pkts_l3_redirect_table_error"].(int)
		ret.So_pkts_l3_redirect_fragmentation_error = in["so_pkts_l3_redirect_fragmentation_error"].(int)
		ret.So_pkts_l3_redirect_table_no_entry_foun = in["so_pkts_l3_redirect_table_no_entry_foun"].(int)
		ret.So_pkts_l3_redirect_invalid_dev_dir = in["so_pkts_l3_redirect_invalid_dev_dir"].(int)
		ret.So_pkts_l3_redirect_chassis_dest_mac_er = in["so_pkts_l3_redirect_chassis_dest_mac_er"].(int)
		ret.So_pkts_l2redirect_vlan_retrieval_error = in["so_pkts_l2redirect_vlan_retrieval_error"].(int)
		ret.So_pkts_l2redirect_port_retrieval_error = in["so_pkts_l2redirect_port_retrieval_error"].(int)
		ret.So_pkts_l2redirect_loop_detect_drop = in["so_pkts_l2redirect_loop_detect_drop"].(int)
		ret.So_pkts_l2redirect_frag_vlan_retrieval_ = in["so_pkts_l2redirect_frag_vlan_retrieval_"].(int)
		ret.So_pkts_l2redirect_tx_frag_vlan_add_fai = in["so_pkts_l2redirect_tx_frag_vlan_add_fai"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPlyrIdGbl2714(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPlyrIdGbl2714 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPlyrIdGbl2714
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsInc2715(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsRate2716(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsInc2715(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsInc2715 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsInc2715
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Total_invalid_playerid_pkts = in["total_invalid_playerid_pkts"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsRate2716(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsRate2716 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsRate2716
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Total_invalid_playerid_pkts = in["total_invalid_playerid_pkts"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSportRate2717(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSportRate2717 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSportRate2717
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSportRateTriggerStatsInc2718(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSportRateTriggerStatsRate2719(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSportRateTriggerStatsInc2718(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSportRateTriggerStatsInc2718 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSportRateTriggerStatsInc2718
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Total_reset = in["total_reset"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSportRateTriggerStatsRate2719(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSportRateTriggerStatsRate2719 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSportRateTriggerStatsRate2719
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Total_reset = in["total_reset"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeLoggingLocalLogGlobal2720(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeLoggingLocalLogGlobal2720 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeLoggingLocalLogGlobal2720
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsInc2721(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsRate2722(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsInc2721(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsInc2721 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsInc2721
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EnqueueFull = in["enqueue_full"].(int)
		ret.EnqueueError = in["enqueue_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsRate2722(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsRate2722 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsRate2722
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.EnqueueFull = in["enqueue_full"].(int)
		ret.EnqueueError = in["enqueue_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlb2723(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlb2723 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlb2723
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlbTriggerStatsInc2724(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlbTriggerStatsRate2725(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlbTriggerStatsInc2724(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlbTriggerStatsInc2724 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlbTriggerStatsInc2724
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mlb_dcmsg_error = in["mlb_dcmsg_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlbTriggerStatsRate2725(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlbTriggerStatsRate2725 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlbTriggerStatsRate2725
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Mlb_dcmsg_error = in["mlb_dcmsg_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbLinkProbe2726(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbLinkProbe2726 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbLinkProbe2726
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsInc2727(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsRate2728(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsInc2727(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsInc2727 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsInc2727
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Err_entry_create_failed = in["err_entry_create_failed"].(int)
		ret.Err_entry_create_oom = in["err_entry_create_oom"].(int)
		ret.Err_entry_insert_failed = in["err_entry_insert_failed"].(int)
		ret.Err_tmpl_probe_create_failed = in["err_tmpl_probe_create_failed"].(int)
		ret.Err_tmpl_probe_create_oom = in["err_tmpl_probe_create_oom"].(int)
		ret.Total_http_response_bad = in["total_http_response_bad"].(int)
		ret.Total_tcp_err = in["total_tcp_err"].(int)
		ret.Err_smart_nat_alloc = in["err_smart_nat_alloc"].(int)
		ret.Err_smart_nat_port_alloc = in["err_smart_nat_port_alloc"].(int)
		ret.Err_l4_sess_alloc = in["err_l4_sess_alloc"].(int)
		ret.Err_probe_tcp_conn_send = in["err_probe_tcp_conn_send"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsRate2728(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsRate2728 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsRate2728
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Err_entry_create_failed = in["err_entry_create_failed"].(int)
		ret.Err_entry_create_oom = in["err_entry_create_oom"].(int)
		ret.Err_entry_insert_failed = in["err_entry_insert_failed"].(int)
		ret.Err_tmpl_probe_create_failed = in["err_tmpl_probe_create_failed"].(int)
		ret.Err_tmpl_probe_create_oom = in["err_tmpl_probe_create_oom"].(int)
		ret.Total_http_response_bad = in["total_http_response_bad"].(int)
		ret.Total_tcp_err = in["total_tcp_err"].(int)
		ret.Err_smart_nat_alloc = in["err_smart_nat_alloc"].(int)
		ret.Err_smart_nat_port_alloc = in["err_smart_nat_port_alloc"].(int)
		ret.Err_l4_sess_alloc = in["err_l4_sess_alloc"].(int)
		ret.Err_probe_tcp_conn_send = in["err_probe_tcp_conn_send"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRpz2729(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRpz2729 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRpz2729
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRpzTriggerStatsInc2730(in["trigger_stats_inc"].([]interface{}))
		ret.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRpzTriggerStatsRate2731(in["trigger_stats_rate"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRpzTriggerStatsInc2730(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRpzTriggerStatsInc2730 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRpzTriggerStatsInc2730
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Set_bw_error = in["set_bw_error"].(int)
		ret.Parse_error = in["parse_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRpzTriggerStatsRate2731(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRpzTriggerStatsRate2731 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRpzTriggerStatsRate2731
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Set_bw_error = in["set_bw_error"].(int)
		ret.Parse_error = in["parse_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsSeverity2732(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsSeverity2732 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsSeverity2732
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplate(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplate {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplate
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.TriggerSysObjStatsChange = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChange2479(d.Get("trigger_sys_obj_stats_change").([]interface{}))
	ret.Inst.TriggerSysObjStatsSeverity = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsSeverity2732(d.Get("trigger_sys_obj_stats_severity").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
