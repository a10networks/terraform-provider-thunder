package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryIpProtoOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_entry_ip_proto_oper`: Operational Status for the object ip-proto\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstEntryIpProtoOperRead,

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
									"src_address_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"port_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"state_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"level_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_connections": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"connection_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_connection_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"connection_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_packet_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"packet_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_kbit_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"kbit_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_frag_packet_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"frag_packet_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat1": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat1_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat2": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat2_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat3": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat3_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat4": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat4_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat5": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat5_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat6": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat6_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat7": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat7_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_app_stat8": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat8_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"age_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"lockup_time_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dynamic_entry_count": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dynamic_entry_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sflow_source_id": {
										Type: schema.TypeString, Optional: true, Description: "",
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
						"all_ports": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"all_src_ports": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"all_ip_protos": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"port_protocol": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"app_stat": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sflow_source_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"hw_blacklisted": {
							Type: schema.TypeString, Optional: true, Description: "",
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
			"port_num": {
				Type: schema.TypeInt, Required: true, Description: "Protocol Number",
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}

func resourceDdosDstEntryIpProtoOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryIpProtoOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryIpProtoOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstEntryIpProtoOperIpFilteringPolicyOper := setObjectDdosDstEntryIpProtoOperIpFilteringPolicyOper(res)
		d.Set("ip_filtering_policy_oper", DdosDstEntryIpProtoOperIpFilteringPolicyOper)
		DdosDstEntryIpProtoOperOper := setObjectDdosDstEntryIpProtoOperOper(res)
		d.Set("oper", DdosDstEntryIpProtoOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstEntryIpProtoOperIpFilteringPolicyOper(ret edpt.DataDdosDstEntryIpProtoOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstEntryIpProtoOperIpFilteringPolicyOperOper(ret.DtDdosDstEntryIpProtoOper.IpFilteringPolicyOper.Oper),
		},
	}
}

func setObjectDdosDstEntryIpProtoOperIpFilteringPolicyOperOper(d edpt.DdosDstEntryIpProtoOperIpFilteringPolicyOperOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceDdosDstEntryIpProtoOperIpFilteringPolicyOperOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceDdosDstEntryIpProtoOperIpFilteringPolicyOperOperRuleList(d []edpt.DdosDstEntryIpProtoOperIpFilteringPolicyOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstEntryIpProtoOperOper(ret edpt.DataDdosDstEntryIpProtoOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ddos_entry_list":         setSliceDdosDstEntryIpProtoOperOperDdos_entry_list(ret.DtDdosDstEntryIpProtoOper.Oper.Ddos_entry_list),
			"entry_displayed_count":   ret.DtDdosDstEntryIpProtoOper.Oper.EntryDisplayedCount,
			"service_displayed_count": ret.DtDdosDstEntryIpProtoOper.Oper.ServiceDisplayedCount,
			"reporting_status":        ret.DtDdosDstEntryIpProtoOper.Oper.ReportingStatus,
			"all_ports":               ret.DtDdosDstEntryIpProtoOper.Oper.AllPorts,
			"all_src_ports":           ret.DtDdosDstEntryIpProtoOper.Oper.AllSrcPorts,
			"all_ip_protos":           ret.DtDdosDstEntryIpProtoOper.Oper.AllIpProtos,
			"port_protocol":           ret.DtDdosDstEntryIpProtoOper.Oper.PortProtocol,
			"app_stat":                ret.DtDdosDstEntryIpProtoOper.Oper.AppStat,
			"sflow_source_id":         ret.DtDdosDstEntryIpProtoOper.Oper.SflowSourceId,
			"hw_blacklisted":          ret.DtDdosDstEntryIpProtoOper.Oper.HwBlacklisted,
			"suffix_request_rate":     ret.DtDdosDstEntryIpProtoOper.Oper.SuffixRequestRate,
			"domain_name":             ret.DtDdosDstEntryIpProtoOper.Oper.DomainName,
		},
	}
}

func setSliceDdosDstEntryIpProtoOperOperDdos_entry_list(d []edpt.DdosDstEntryIpProtoOperOperDdos_entry_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dst_address_str"] = item.DstAddressStr
		in["src_address_str"] = item.SrcAddressStr
		in["port_str"] = item.PortStr
		in["state_str"] = item.StateStr
		in["level_str"] = item.LevelStr
		in["current_connections"] = item.CurrentConnections
		in["connection_limit"] = item.ConnectionLimit
		in["current_connection_rate"] = item.CurrentConnectionRate
		in["connection_rate_limit"] = item.ConnectionRateLimit
		in["current_packet_rate"] = item.CurrentPacketRate
		in["packet_rate_limit"] = item.PacketRateLimit
		in["current_kbit_rate"] = item.CurrentKbitRate
		in["kbit_rate_limit"] = item.KbitRateLimit
		in["current_frag_packet_rate"] = item.CurrentFragPacketRate
		in["frag_packet_rate_limit"] = item.FragPacketRateLimit
		in["current_app_stat1"] = item.CurrentAppStat1
		in["app_stat1_limit"] = item.AppStat1Limit
		in["current_app_stat2"] = item.CurrentAppStat2
		in["app_stat2_limit"] = item.AppStat2Limit
		in["current_app_stat3"] = item.CurrentAppStat3
		in["app_stat3_limit"] = item.AppStat3Limit
		in["current_app_stat4"] = item.CurrentAppStat4
		in["app_stat4_limit"] = item.AppStat4Limit
		in["current_app_stat5"] = item.CurrentAppStat5
		in["app_stat5_limit"] = item.AppStat5Limit
		in["current_app_stat6"] = item.CurrentAppStat6
		in["app_stat6_limit"] = item.AppStat6Limit
		in["current_app_stat7"] = item.CurrentAppStat7
		in["app_stat7_limit"] = item.AppStat7Limit
		in["current_app_stat8"] = item.CurrentAppStat8
		in["app_stat8_limit"] = item.AppStat8Limit
		in["age_str"] = item.AgeStr
		in["lockup_time_str"] = item.LockupTimeStr
		in["dynamic_entry_count"] = item.DynamicEntryCount
		in["dynamic_entry_limit"] = item.DynamicEntryLimit
		in["sflow_source_id"] = item.SflowSourceId
		in["debug_str"] = item.DebugStr
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstEntryIpProtoOperIpFilteringPolicyOper(d []interface{}) edpt.DdosDstEntryIpProtoOperIpFilteringPolicyOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryIpProtoOperIpFilteringPolicyOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstEntryIpProtoOperIpFilteringPolicyOperOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryIpProtoOperIpFilteringPolicyOperOper(d []interface{}) edpt.DdosDstEntryIpProtoOperIpFilteringPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryIpProtoOperIpFilteringPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstEntryIpProtoOperIpFilteringPolicyOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryIpProtoOperIpFilteringPolicyOperOperRuleList(d []interface{}) []edpt.DdosDstEntryIpProtoOperIpFilteringPolicyOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryIpProtoOperIpFilteringPolicyOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryIpProtoOperIpFilteringPolicyOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntryIpProtoOperOper(d []interface{}) edpt.DdosDstEntryIpProtoOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryIpProtoOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ddos_entry_list = getSliceDdosDstEntryIpProtoOperOperDdos_entry_list(in["ddos_entry_list"].([]interface{}))
		ret.EntryDisplayedCount = in["entry_displayed_count"].(int)
		ret.ServiceDisplayedCount = in["service_displayed_count"].(int)
		ret.ReportingStatus = in["reporting_status"].(int)
		ret.AllPorts = in["all_ports"].(int)
		ret.AllSrcPorts = in["all_src_ports"].(int)
		ret.AllIpProtos = in["all_ip_protos"].(int)
		ret.PortProtocol = in["port_protocol"].(string)
		ret.AppStat = in["app_stat"].(int)
		ret.SflowSourceId = in["sflow_source_id"].(int)
		ret.HwBlacklisted = in["hw_blacklisted"].(string)
		ret.SuffixRequestRate = in["suffix_request_rate"].(int)
		ret.DomainName = in["domain_name"].(string)
	}
	return ret
}

func getSliceDdosDstEntryIpProtoOperOperDdos_entry_list(d []interface{}) []edpt.DdosDstEntryIpProtoOperOperDdos_entry_list {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryIpProtoOperOperDdos_entry_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryIpProtoOperOperDdos_entry_list
		oi.DstAddressStr = in["dst_address_str"].(string)
		oi.SrcAddressStr = in["src_address_str"].(string)
		oi.PortStr = in["port_str"].(string)
		oi.StateStr = in["state_str"].(string)
		oi.LevelStr = in["level_str"].(string)
		oi.CurrentConnections = in["current_connections"].(string)
		oi.ConnectionLimit = in["connection_limit"].(string)
		oi.CurrentConnectionRate = in["current_connection_rate"].(string)
		oi.ConnectionRateLimit = in["connection_rate_limit"].(string)
		oi.CurrentPacketRate = in["current_packet_rate"].(string)
		oi.PacketRateLimit = in["packet_rate_limit"].(string)
		oi.CurrentKbitRate = in["current_kbit_rate"].(string)
		oi.KbitRateLimit = in["kbit_rate_limit"].(string)
		oi.CurrentFragPacketRate = in["current_frag_packet_rate"].(string)
		oi.FragPacketRateLimit = in["frag_packet_rate_limit"].(string)
		oi.CurrentAppStat1 = in["current_app_stat1"].(string)
		oi.AppStat1Limit = in["app_stat1_limit"].(string)
		oi.CurrentAppStat2 = in["current_app_stat2"].(string)
		oi.AppStat2Limit = in["app_stat2_limit"].(string)
		oi.CurrentAppStat3 = in["current_app_stat3"].(string)
		oi.AppStat3Limit = in["app_stat3_limit"].(string)
		oi.CurrentAppStat4 = in["current_app_stat4"].(string)
		oi.AppStat4Limit = in["app_stat4_limit"].(string)
		oi.CurrentAppStat5 = in["current_app_stat5"].(string)
		oi.AppStat5Limit = in["app_stat5_limit"].(string)
		oi.CurrentAppStat6 = in["current_app_stat6"].(string)
		oi.AppStat6Limit = in["app_stat6_limit"].(string)
		oi.CurrentAppStat7 = in["current_app_stat7"].(string)
		oi.AppStat7Limit = in["app_stat7_limit"].(string)
		oi.CurrentAppStat8 = in["current_app_stat8"].(string)
		oi.AppStat8Limit = in["app_stat8_limit"].(string)
		oi.AgeStr = in["age_str"].(string)
		oi.LockupTimeStr = in["lockup_time_str"].(string)
		oi.DynamicEntryCount = in["dynamic_entry_count"].(string)
		oi.DynamicEntryLimit = in["dynamic_entry_limit"].(string)
		oi.SflowSourceId = in["sflow_source_id"].(string)
		oi.DebugStr = in["debug_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstEntryIpProtoOper(d *schema.ResourceData) edpt.DdosDstEntryIpProtoOper {
	var ret edpt.DdosDstEntryIpProtoOper

	ret.IpFilteringPolicyOper = getObjectDdosDstEntryIpProtoOperIpFilteringPolicyOper(d.Get("ip_filtering_policy_oper").([]interface{}))

	ret.Oper = getObjectDdosDstEntryIpProtoOperOper(d.Get("oper").([]interface{}))

	ret.PortNum = d.Get("port_num").(int)

	ret.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
