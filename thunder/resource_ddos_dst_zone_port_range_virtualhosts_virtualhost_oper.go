package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortRangeVirtualhostsVirtualhostOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_range_virtualhosts_virtualhost_oper`: Operational Status for the object virtualhost\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortRangeVirtualhostsVirtualhostOperRead,

		Schema: map[string]*schema.Schema{
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
									"current_app_stat1": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_app_stat1_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"app_stat1_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat2": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_app_stat2_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"app_stat2_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat3": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_app_stat3_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"app_stat3_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat4": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_app_stat4_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"app_stat4_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat5": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_app_stat5_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"app_stat5_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat6": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_app_stat6_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"app_stat6_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat7": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_app_stat7_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"app_stat7_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat8": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_app_stat8_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"app_stat8_limit": {
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
									"dynamic_entry_warn_state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sflow_source_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http_filter_rates": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"http_filter_rate_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"is_http_filter_rate_limit_exceed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"current_http_filter_rate": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"http_filter_rate_limit": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"response_size_rates": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"response_size_rate_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"is_response_size_rate_limit_exceed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"current_response_size_rate": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"response_size_rate_limit": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
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
						"l4_ext_rate": {
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
			"vhost": {
				Type: schema.TypeString, Required: true, Description: "name for virtualhost",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"port_range_start": {
				Type: schema.TypeString, Required: true, Description: "PortRangeStart",
			},
			"port_range_end": {
				Type: schema.TypeString, Required: true, Description: "PortRangeEnd",
			},
		},
	}
}

func resourceDdosDstZonePortRangeVirtualhostsVirtualhostOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeVirtualhostsVirtualhostOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeVirtualhostsVirtualhostOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortRangeVirtualhostsVirtualhostOperOper := setObjectDdosDstZonePortRangeVirtualhostsVirtualhostOperOper(res)
		d.Set("oper", DdosDstZonePortRangeVirtualhostsVirtualhostOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortRangeVirtualhostsVirtualhostOperOper(ret edpt.DataDdosDstZonePortRangeVirtualhostsVirtualhostOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ddos_entry_list":         setSliceDdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_list(ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.Ddos_entry_list),
			"entry_displayed_count":   ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.EntryDisplayedCount,
			"service_displayed_count": ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.ServiceDisplayedCount,
			"reporting_status":        ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.ReportingStatus,
			"sources":                 ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.Sources,
			"overflow_policy":         ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.OverflowPolicy,
			"sources_all_entries":     ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.SourcesAllEntries,
			"class_list":              ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.ClassList,
			"subnet_ip_addr":          ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.SubnetIpAddr,
			"subnet_ipv6_addr":        ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.SubnetIpv6Addr,
			"ipv6":                    ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.Ipv6,
			"exceeded":                ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.Exceeded,
			"black_listed":            ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.BlackListed,
			"white_listed":            ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.WhiteListed,
			"authenticated":           ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.Authenticated,
			"level":                   ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.Level,
			"app_stat":                ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.AppStat,
			"indicators":              ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.Indicators,
			"indicator_detail":        ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.IndicatorDetail,
			"l4_ext_rate":             ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.L4ExtRate,
			"hw_blacklisted":          ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.HwBlacklisted,
			"suffix_request_rate":     ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.SuffixRequestRate,
			"domain_name":             ret.DtDdosDstZonePortRangeVirtualhostsVirtualhostOper.Oper.DomainName,
		},
	}
}

func setSliceDdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_list(d []edpt.DdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_list) []map[string]interface{} {
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
		in["current_app_stat1"] = item.CurrentAppStat1
		in["is_app_stat1_exceed"] = item.IsAppStat1Exceed
		in["app_stat1_limit"] = item.AppStat1Limit
		in["current_app_stat2"] = item.CurrentAppStat2
		in["is_app_stat2_exceed"] = item.IsAppStat2Exceed
		in["app_stat2_limit"] = item.AppStat2Limit
		in["current_app_stat3"] = item.CurrentAppStat3
		in["is_app_stat3_exceed"] = item.IsAppStat3Exceed
		in["app_stat3_limit"] = item.AppStat3Limit
		in["current_app_stat4"] = item.CurrentAppStat4
		in["is_app_stat4_exceed"] = item.IsAppStat4Exceed
		in["app_stat4_limit"] = item.AppStat4Limit
		in["current_app_stat5"] = item.CurrentAppStat5
		in["is_app_stat5_exceed"] = item.IsAppStat5Exceed
		in["app_stat5_limit"] = item.AppStat5Limit
		in["current_app_stat6"] = item.CurrentAppStat6
		in["is_app_stat6_exceed"] = item.IsAppStat6Exceed
		in["app_stat6_limit"] = item.AppStat6Limit
		in["current_app_stat7"] = item.CurrentAppStat7
		in["is_app_stat7_exceed"] = item.IsAppStat7Exceed
		in["app_stat7_limit"] = item.AppStat7Limit
		in["current_app_stat8"] = item.CurrentAppStat8
		in["is_app_stat8_exceed"] = item.IsAppStat8Exceed
		in["app_stat8_limit"] = item.AppStat8Limit
		in["age"] = item.Age
		in["lockup_time"] = item.LockupTime
		in["dynamic_entry_count"] = item.DynamicEntryCount
		in["dynamic_entry_limit"] = item.DynamicEntryLimit
		in["dynamic_entry_warn_state"] = item.DynamicEntryWarnState
		in["sflow_source_id"] = item.SflowSourceId
		in["http_filter_rates"] = setSliceDdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_listHttpFilterRates(item.HttpFilterRates)
		in["response_size_rates"] = setSliceDdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_listResponseSizeRates(item.ResponseSizeRates)
		in["debug_str"] = item.DebugStr
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_listHttpFilterRates(d []edpt.DdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_listHttpFilterRates) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["http_filter_rate_name"] = item.HttpFilterRateName
		in["is_http_filter_rate_limit_exceed"] = item.IsHttpFilterRateLimitExceed
		in["current_http_filter_rate"] = item.CurrentHttpFilterRate
		in["http_filter_rate_limit"] = item.HttpFilterRateLimit
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_listResponseSizeRates(d []edpt.DdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_listResponseSizeRates) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["response_size_rate_name"] = item.ResponseSizeRateName
		in["is_response_size_rate_limit_exceed"] = item.IsResponseSizeRateLimitExceed
		in["current_response_size_rate"] = item.CurrentResponseSizeRate
		in["response_size_rate_limit"] = item.ResponseSizeRateLimit
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZonePortRangeVirtualhostsVirtualhostOperOper(d []interface{}) edpt.DdosDstZonePortRangeVirtualhostsVirtualhostOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeVirtualhostsVirtualhostOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
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
		ret.L4ExtRate = in["l4_ext_rate"].(int)
		ret.HwBlacklisted = in["hw_blacklisted"].(int)
		ret.SuffixRequestRate = in["suffix_request_rate"].(int)
		ret.DomainName = in["domain_name"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_list(d []interface{}) []edpt.DdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_list
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
		oi.CurrentAppStat1 = in["current_app_stat1"].(string)
		oi.IsAppStat1Exceed = in["is_app_stat1_exceed"].(int)
		oi.AppStat1Limit = in["app_stat1_limit"].(string)
		oi.CurrentAppStat2 = in["current_app_stat2"].(string)
		oi.IsAppStat2Exceed = in["is_app_stat2_exceed"].(int)
		oi.AppStat2Limit = in["app_stat2_limit"].(string)
		oi.CurrentAppStat3 = in["current_app_stat3"].(string)
		oi.IsAppStat3Exceed = in["is_app_stat3_exceed"].(int)
		oi.AppStat3Limit = in["app_stat3_limit"].(string)
		oi.CurrentAppStat4 = in["current_app_stat4"].(string)
		oi.IsAppStat4Exceed = in["is_app_stat4_exceed"].(int)
		oi.AppStat4Limit = in["app_stat4_limit"].(string)
		oi.CurrentAppStat5 = in["current_app_stat5"].(string)
		oi.IsAppStat5Exceed = in["is_app_stat5_exceed"].(int)
		oi.AppStat5Limit = in["app_stat5_limit"].(string)
		oi.CurrentAppStat6 = in["current_app_stat6"].(string)
		oi.IsAppStat6Exceed = in["is_app_stat6_exceed"].(int)
		oi.AppStat6Limit = in["app_stat6_limit"].(string)
		oi.CurrentAppStat7 = in["current_app_stat7"].(string)
		oi.IsAppStat7Exceed = in["is_app_stat7_exceed"].(int)
		oi.AppStat7Limit = in["app_stat7_limit"].(string)
		oi.CurrentAppStat8 = in["current_app_stat8"].(string)
		oi.IsAppStat8Exceed = in["is_app_stat8_exceed"].(int)
		oi.AppStat8Limit = in["app_stat8_limit"].(string)
		oi.Age = in["age"].(int)
		oi.LockupTime = in["lockup_time"].(int)
		oi.DynamicEntryCount = in["dynamic_entry_count"].(string)
		oi.DynamicEntryLimit = in["dynamic_entry_limit"].(string)
		oi.DynamicEntryWarnState = in["dynamic_entry_warn_state"].(string)
		oi.SflowSourceId = in["sflow_source_id"].(int)
		oi.HttpFilterRates = getSliceDdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_listHttpFilterRates(in["http_filter_rates"].([]interface{}))
		oi.ResponseSizeRates = getSliceDdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_listResponseSizeRates(in["response_size_rates"].([]interface{}))
		oi.DebugStr = in["debug_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_listHttpFilterRates(d []interface{}) []edpt.DdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_listHttpFilterRates {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_listHttpFilterRates, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_listHttpFilterRates
		oi.HttpFilterRateName = in["http_filter_rate_name"].(string)
		oi.IsHttpFilterRateLimitExceed = in["is_http_filter_rate_limit_exceed"].(int)
		oi.CurrentHttpFilterRate = in["current_http_filter_rate"].(string)
		oi.HttpFilterRateLimit = in["http_filter_rate_limit"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_listResponseSizeRates(d []interface{}) []edpt.DdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_listResponseSizeRates {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_listResponseSizeRates, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeVirtualhostsVirtualhostOperOperDdos_entry_listResponseSizeRates
		oi.ResponseSizeRateName = in["response_size_rate_name"].(string)
		oi.IsResponseSizeRateLimitExceed = in["is_response_size_rate_limit_exceed"].(int)
		oi.CurrentResponseSizeRate = in["current_response_size_rate"].(string)
		oi.ResponseSizeRateLimit = in["response_size_rate_limit"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortRangeVirtualhostsVirtualhostOper(d *schema.ResourceData) edpt.DdosDstZonePortRangeVirtualhostsVirtualhostOper {
	var ret edpt.DdosDstZonePortRangeVirtualhostsVirtualhostOper

	ret.Oper = getObjectDdosDstZonePortRangeVirtualhostsVirtualhostOperOper(d.Get("oper").([]interface{}))

	ret.Vhost = d.Get("vhost").(string)

	ret.Protocol = d.Get("protocol").(string)

	ret.ZoneName = d.Get("zone_name").(string)

	ret.PortRangeStart = d.Get("port_range_start").(string)

	ret.PortRangeEnd = d.Get("port_range_end").(string)
	return ret
}
