package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneSrcPortZoneSrcPortOtherOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_src_port_zone_src_port_other_oper`: Operational Status for the object zone-src-port-other\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneSrcPortZoneSrcPortOtherOperRead,

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
						"sources_all_entries": {
							Type: schema.TypeInt, Optional: true, Description: "",
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
						"hw_blacklisted": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"port_ind": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"indicators": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"indicator_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"indicator_index": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"rate": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"indicator_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"level": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"zone_threshold": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"source_threshold": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
											},
										},
									},
									"detection_data_source": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_level": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"details": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"port_other": {
				Type: schema.TypeString, Required: true, Description: "'other': other;",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'udp': UDP port; 'tcp': TCP Port;",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}

func resourceDdosDstZoneSrcPortZoneSrcPortOtherOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortZoneSrcPortOtherOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortZoneSrcPortOtherOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneSrcPortZoneSrcPortOtherOperOper := setObjectDdosDstZoneSrcPortZoneSrcPortOtherOperOper(res)
		d.Set("oper", DdosDstZoneSrcPortZoneSrcPortOtherOperOper)
		DdosDstZoneSrcPortZoneSrcPortOtherOperPortInd := setObjectDdosDstZoneSrcPortZoneSrcPortOtherOperPortInd(res)
		d.Set("port_ind", DdosDstZoneSrcPortZoneSrcPortOtherOperPortInd)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneSrcPortZoneSrcPortOtherOperOper(ret edpt.DataDdosDstZoneSrcPortZoneSrcPortOtherOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ddos_entry_list":         setSliceDdosDstZoneSrcPortZoneSrcPortOtherOperOperDdos_entry_list(ret.DtDdosDstZoneSrcPortZoneSrcPortOtherOper.Oper.Ddos_entry_list),
			"entry_displayed_count":   ret.DtDdosDstZoneSrcPortZoneSrcPortOtherOper.Oper.EntryDisplayedCount,
			"service_displayed_count": ret.DtDdosDstZoneSrcPortZoneSrcPortOtherOper.Oper.ServiceDisplayedCount,
			"reporting_status":        ret.DtDdosDstZoneSrcPortZoneSrcPortOtherOper.Oper.ReportingStatus,
			"sources":                 ret.DtDdosDstZoneSrcPortZoneSrcPortOtherOper.Oper.Sources,
			"sources_all_entries":     ret.DtDdosDstZoneSrcPortZoneSrcPortOtherOper.Oper.SourcesAllEntries,
			"subnet_ip_addr":          ret.DtDdosDstZoneSrcPortZoneSrcPortOtherOper.Oper.SubnetIpAddr,
			"subnet_ipv6_addr":        ret.DtDdosDstZoneSrcPortZoneSrcPortOtherOper.Oper.SubnetIpv6Addr,
			"ipv6":                    ret.DtDdosDstZoneSrcPortZoneSrcPortOtherOper.Oper.Ipv6,
			"hw_blacklisted":          ret.DtDdosDstZoneSrcPortZoneSrcPortOtherOper.Oper.HwBlacklisted,
		},
	}
}

func setSliceDdosDstZoneSrcPortZoneSrcPortOtherOperOperDdos_entry_list(d []edpt.DdosDstZoneSrcPortZoneSrcPortOtherOperOperDdos_entry_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dst_address_str"] = item.DstAddressStr
		in["bw_state"] = item.BwState
		in["is_auth_passed"] = item.Is_auth_passed
		in["level"] = item.Level
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
		in["sflow_source_id"] = item.SflowSourceId
		in["debug_str"] = item.DebugStr
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneSrcPortZoneSrcPortOtherOperPortInd(ret edpt.DataDdosDstZoneSrcPortZoneSrcPortOtherOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOper(ret.DtDdosDstZoneSrcPortZoneSrcPortOtherOper.PortInd.Oper),
		},
	}
}

func setObjectDdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOper(d edpt.DdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicators(d.Indicators)

	in["detection_data_source"] = d.DetectionDataSource

	in["current_level"] = d.CurrentLevel

	in["details"] = d.Details
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicators(d []edpt.DdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["rate"] = item.Rate
		in["indicator_cfg"] = setSliceDdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicatorsIndicatorCfg(item.IndicatorCfg)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicatorsIndicatorCfg(d []edpt.DdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicatorsIndicatorCfg) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["level"] = item.Level
		in["zone_threshold"] = item.ZoneThreshold
		in["source_threshold"] = item.SourceThreshold
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZoneSrcPortZoneSrcPortOtherOperOper(d []interface{}) edpt.DdosDstZoneSrcPortZoneSrcPortOtherOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortZoneSrcPortOtherOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstZoneSrcPortZoneSrcPortOtherOperOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
		ret.EntryDisplayedCount = in["entry_displayed_count"].(int)
		ret.ServiceDisplayedCount = in["service_displayed_count"].(int)
		ret.ReportingStatus = in["reporting_status"].(int)
		ret.Sources = in["sources"].(int)
		ret.SourcesAllEntries = in["sources_all_entries"].(int)
		ret.SubnetIpAddr = in["subnet_ip_addr"].(string)
		ret.SubnetIpv6Addr = in["subnet_ipv6_addr"].(string)
		ret.Ipv6 = in["ipv6"].(string)
		ret.HwBlacklisted = in["hw_blacklisted"].(int)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortZoneSrcPortOtherOperOperDdos_entry_list(d []interface{}) []edpt.DdosDstZoneSrcPortZoneSrcPortOtherOperOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortZoneSrcPortOtherOperOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortZoneSrcPortOtherOperOperDdos_entry_list
		oi.DstAddressStr = in["dst_address_str"].(string)
		oi.BwState = in["bw_state"].(string)
		oi.Is_auth_passed = in["is_auth_passed"].(string)
		oi.Level = in["level"].(int)
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
		oi.SflowSourceId = in["sflow_source_id"].(int)
		oi.DebugStr = in["debug_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneSrcPortZoneSrcPortOtherOperPortInd(d []interface{}) edpt.DdosDstZoneSrcPortZoneSrcPortOtherOperPortInd {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortZoneSrcPortOtherOperPortInd
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOper(d []interface{}) edpt.DdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicators(in["indicators"].([]interface{}))
		ret.DetectionDataSource = in["detection_data_source"].(string)
		ret.CurrentLevel = in["current_level"].(string)
		ret.Details = in["details"].(int)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicators(d []interface{}) []edpt.DdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.IndicatorCfg = getSliceDdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicatorsIndicatorCfg(in["indicator_cfg"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicatorsIndicatorCfg(d []interface{}) []edpt.DdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicatorsIndicatorCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicatorsIndicatorCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicatorsIndicatorCfg
		oi.Level = in["level"].(int)
		oi.ZoneThreshold = in["zone_threshold"].(string)
		oi.SourceThreshold = in["source_threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneSrcPortZoneSrcPortOtherOper(d *schema.ResourceData) edpt.DdosDstZoneSrcPortZoneSrcPortOtherOper {
	var ret edpt.DdosDstZoneSrcPortZoneSrcPortOtherOper

	ret.Oper = getObjectDdosDstZoneSrcPortZoneSrcPortOtherOperOper(d.Get("oper").([]interface{}))

	ret.PortInd = getObjectDdosDstZoneSrcPortZoneSrcPortOtherOperPortInd(d.Get("port_ind").([]interface{}))

	ret.PortOther = d.Get("port_other").(string)

	ret.Protocol = d.Get("protocol").(string)

	ret.ZoneName = d.Get("zone_name").(string)
	return ret
}
