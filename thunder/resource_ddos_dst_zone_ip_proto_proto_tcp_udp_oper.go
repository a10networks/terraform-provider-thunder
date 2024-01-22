package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoTcpUdpOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_ip_proto_proto_tcp_udp_oper`: Operational Status for the object proto-tcp-udp\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneIpProtoProtoTcpUdpOperRead,

		Schema: map[string]*schema.Schema{
			"ip_filtering_policy_oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rule_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"seq": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"hits": {
													Type: schema.TypeInt, Optional: true, Description: "",
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
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ddos_entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dst_address_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"bw_state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_auth_passed": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"level": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bl_reasoning_rcode": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"bl_reasoning_timestamp": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_connections": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_connections_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connection_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_connection_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_connection_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connection_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_packet_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_packet_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"packet_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_kbit_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_kbit_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"kbit_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_frag_packet_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_frag_packet_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"frag_packet_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"age": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lockup_time": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dynamic_entry_count": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dynamic_entry_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sflow_source_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"debug_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"entry_displayed_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"service_displayed_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"reporting_status": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sources": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"overflow_policy": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sources_all_entries": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"class_list": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"subnet_ip_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"subnet_ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"black_listed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"white_listed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"authenticated": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"level": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"app_stat": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"indicators": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"indicator_detail": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"hw_blacklisted": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"suffix_request_rate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"domain_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': ip-proto tcp; 'udp': ip-proto udp;",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}

func resourceDdosDstZoneIpProtoProtoTcpUdpOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoTcpUdpOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoTcpUdpOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOper := setObjectDdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOper(res)
		d.Set("ip_filtering_policy_oper", DdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOper)
		DdosDstZoneIpProtoProtoTcpUdpOperOper := setObjectDdosDstZoneIpProtoProtoTcpUdpOperOper(res)
		d.Set("oper", DdosDstZoneIpProtoProtoTcpUdpOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOper(ret edpt.DataDdosDstZoneIpProtoProtoTcpUdpOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOperOper(ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.IpFilteringPolicyOper.Oper),
		},
	}
}

func setObjectDdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOperOper(d edpt.DdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOperOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceDdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOperOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOperOperRuleList(d []edpt.DdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneIpProtoProtoTcpUdpOperOper(ret edpt.DataDdosDstZoneIpProtoProtoTcpUdpOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ddos_entry_list":         setSliceDdosDstZoneIpProtoProtoTcpUdpOperOperDdos_entry_list(ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.Ddos_entry_list),
			"entry_displayed_count":   ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.EntryDisplayedCount,
			"service_displayed_count": ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.ServiceDisplayedCount,
			"reporting_status":        ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.ReportingStatus,
			"sources":                 ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.Sources,
			"overflow_policy":         ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.OverflowPolicy,
			"sources_all_entries":     ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.SourcesAllEntries,
			"class_list":              ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.ClassList,
			"subnet_ip_addr":          ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.SubnetIpAddr,
			"subnet_ipv6_addr":        ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.SubnetIpv6Addr,
			"ipv6":                    ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.Ipv6,
			"exceeded":                ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.Exceeded,
			"black_listed":            ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.BlackListed,
			"white_listed":            ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.WhiteListed,
			"authenticated":           ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.Authenticated,
			"level":                   ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.Level,
			"app_stat":                ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.AppStat,
			"indicators":              ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.Indicators,
			"indicator_detail":        ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.IndicatorDetail,
			"hw_blacklisted":          ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.HwBlacklisted,
			"suffix_request_rate":     ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.SuffixRequestRate,
			"domain_name":             ret.DtDdosDstZoneIpProtoProtoTcpUdpOper.Oper.DomainName,
		},
	}
}

func setSliceDdosDstZoneIpProtoProtoTcpUdpOperOperDdos_entry_list(d []edpt.DdosDstZoneIpProtoProtoTcpUdpOperOperDdos_entry_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dst_address_str"] = item.DstAddressStr
		in["bw_state"] = item.BwState
		in["is_auth_passed"] = item.Is_auth_passed
		in["level"] = item.Level
		in["bl_reasoning_rcode"] = item.BlReasoningRcode
		in["bl_reasoning_timestamp"] = item.BlReasoningTimestamp
		in["current_connections"] = item.CurrentConnections
		in["is_connections_exceed"] = item.IsConnectionsExceed
		in["connection_limit"] = item.ConnectionLimit
		in["current_connection_rate"] = item.CurrentConnectionRate
		in["is_connection_rate_exceed"] = item.IsConnectionRateExceed
		in["connection_rate_limit"] = item.ConnectionRateLimit
		in["current_packet_rate"] = item.CurrentPacketRate
		in["is_packet_rate_exceed"] = item.IsPacketRateExceed
		in["packet_rate_limit"] = item.PacketRateLimit
		in["current_kbit_rate"] = item.CurrentKbitRate
		in["is_kbit_rate_exceed"] = item.IsKbitRateExceed
		in["kbit_rate_limit"] = item.KbitRateLimit
		in["current_frag_packet_rate"] = item.CurrentFragPacketRate
		in["is_frag_packet_rate_exceed"] = item.IsFragPacketRateExceed
		in["frag_packet_rate_limit"] = item.FragPacketRateLimit
		in["age"] = item.Age
		in["lockup_time"] = item.LockupTime
		in["dynamic_entry_count"] = item.DynamicEntryCount
		in["dynamic_entry_limit"] = item.DynamicEntryLimit
		in["sflow_source_id"] = item.SflowSourceId
		in["debug_str"] = item.DebugStr
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOperOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOperOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOperOperRuleList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoTcpUdpOperOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoTcpUdpOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoTcpUdpOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstZoneIpProtoProtoTcpUdpOperOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
		ret.EntryDisplayedCount = in["entry_displayed_count"].(int)
		ret.ServiceDisplayedCount = in["service_displayed_count"].(int)
		ret.ReportingStatus = in["reporting_status"].(int)
		ret.Sources = in["sources"].(int)
		ret.OverflowPolicy = in["overflow_policy"].(int)
		ret.SourcesAllEntries = in["sources_all_entries"].(int)
		ret.ClassList = in["class_list"].(string)
		ret.SubnetIpAddr = in["subnet_ip_addr"].(string)
		ret.SubnetIpv6Addr = in["subnet_ipv6_addr"].(string)
		ret.Ipv6 = in["ipv6"].(string)
		ret.Exceeded = in["exceeded"].(int)
		ret.BlackListed = in["black_listed"].(int)
		ret.WhiteListed = in["white_listed"].(int)
		ret.Authenticated = in["authenticated"].(int)
		ret.Level = in["level"].(int)
		ret.AppStat = in["app_stat"].(int)
		ret.Indicators = in["indicators"].(int)
		ret.IndicatorDetail = in["indicator_detail"].(int)
		ret.HwBlacklisted = in["hw_blacklisted"].(int)
		ret.SuffixRequestRate = in["suffix_request_rate"].(int)
		ret.DomainName = in["domain_name"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoTcpUdpOperOperDdos_entry_list(d []interface{}) []edpt.DdosDstZoneIpProtoProtoTcpUdpOperOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoTcpUdpOperOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoTcpUdpOperOperDdos_entry_list
		oi.DstAddressStr = in["dst_address_str"].(string)
		oi.BwState = in["bw_state"].(string)
		oi.Is_auth_passed = in["is_auth_passed"].(string)
		oi.Level = in["level"].(int)
		oi.BlReasoningRcode = in["bl_reasoning_rcode"].(string)
		oi.BlReasoningTimestamp = in["bl_reasoning_timestamp"].(string)
		oi.CurrentConnections = in["current_connections"].(string)
		oi.IsConnectionsExceed = in["is_connections_exceed"].(int)
		oi.ConnectionLimit = in["connection_limit"].(string)
		oi.CurrentConnectionRate = in["current_connection_rate"].(string)
		oi.IsConnectionRateExceed = in["is_connection_rate_exceed"].(int)
		oi.ConnectionRateLimit = in["connection_rate_limit"].(string)
		oi.CurrentPacketRate = in["current_packet_rate"].(string)
		oi.IsPacketRateExceed = in["is_packet_rate_exceed"].(int)
		oi.PacketRateLimit = in["packet_rate_limit"].(string)
		oi.CurrentKbitRate = in["current_kbit_rate"].(string)
		oi.IsKbitRateExceed = in["is_kbit_rate_exceed"].(int)
		oi.KbitRateLimit = in["kbit_rate_limit"].(string)
		oi.CurrentFragPacketRate = in["current_frag_packet_rate"].(string)
		oi.IsFragPacketRateExceed = in["is_frag_packet_rate_exceed"].(int)
		oi.FragPacketRateLimit = in["frag_packet_rate_limit"].(string)
		oi.Age = in["age"].(int)
		oi.LockupTime = in["lockup_time"].(int)
		oi.DynamicEntryCount = in["dynamic_entry_count"].(string)
		oi.DynamicEntryLimit = in["dynamic_entry_limit"].(string)
		oi.SflowSourceId = in["sflow_source_id"].(int)
		oi.DebugStr = in["debug_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoTcpUdpOper(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoTcpUdpOper {
	var ret edpt.DdosDstZoneIpProtoProtoTcpUdpOper

	ret.IpFilteringPolicyOper = getObjectDdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOper(d.Get("ip_filtering_policy_oper").([]interface{}))

	ret.Oper = getObjectDdosDstZoneIpProtoProtoTcpUdpOperOper(d.Get("oper").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.ZoneName = d.Get("zone_name").(string)
	return ret
}
