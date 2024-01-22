package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDnsCacheOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dns_cache_oper`: Operational Status for the object dns-cache\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDnsCacheOperRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "DNS Cache Instance Name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_entries": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fqdn_name": {
										Type: schema.TypeString, Optional: true, Description: "FQDN Name",
									},
									"fqdn_manual_override_action": {
										Type: schema.TypeString, Optional: true, Description: "FQDN Manual Override Action",
									},
									"wild_card_node": {
										Type: schema.TypeString, Optional: true, Description: "Wild Card Node",
									},
									"delegation_node": {
										Type: schema.TypeString, Optional: true, Description: "Delegation Node",
									},
									"empty_non_terminal_node": {
										Type: schema.TypeString, Optional: true, Description: "Empty Non-Terminal Node",
									},
									"record_types": {
										Type: schema.TypeString, Optional: true, Description: "Type of the records supported",
									},
								},
							},
						},
						"response_status": {
							Type: schema.TypeString, Optional: true, Description: "response status",
						},
						"response_flag": {
							Type: schema.TypeString, Optional: true, Description: "response flag",
						},
						"answer_section_record_count": {
							Type: schema.TypeInt, Optional: true, Description: "Answer section record Count",
						},
						"answer_section_size": {
							Type: schema.TypeInt, Optional: true, Description: "Answer section size",
						},
						"authority_section_record_count": {
							Type: schema.TypeInt, Optional: true, Description: "Authority section record Count",
						},
						"authority_section_size": {
							Type: schema.TypeInt, Optional: true, Description: "Autority section size",
						},
						"additional_section_record_count": {
							Type: schema.TypeInt, Optional: true, Description: "Additional section record Count",
						},
						"additional_section_size": {
							Type: schema.TypeInt, Optional: true, Description: "Additional section size",
						},
						"answer_section": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"record_domain_name": {
										Type: schema.TypeString, Optional: true, Description: "Name of the Domain",
									},
									"record_type": {
										Type: schema.TypeString, Optional: true, Description: "Type of the record",
									},
									"record_class": {
										Type: schema.TypeString, Optional: true, Description: "Class",
									},
									"record_ttl": {
										Type: schema.TypeInt, Optional: true, Description: "TTL",
									},
									"record_data": {
										Type: schema.TypeString, Optional: true, Description: "Record Data",
									},
								},
							},
						},
						"authoritative_section": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"record_domain_name": {
										Type: schema.TypeString, Optional: true, Description: "Name of the Domain",
									},
									"record_type": {
										Type: schema.TypeString, Optional: true, Description: "Type of the record",
									},
									"record_class": {
										Type: schema.TypeString, Optional: true, Description: "Class",
									},
									"record_ttl": {
										Type: schema.TypeInt, Optional: true, Description: "TTL",
									},
									"record_data": {
										Type: schema.TypeString, Optional: true, Description: "Record Data",
									},
								},
							},
						},
						"additional_section": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"record_domain_name": {
										Type: schema.TypeString, Optional: true, Description: "Name of the Domain",
									},
									"record_type": {
										Type: schema.TypeString, Optional: true, Description: "Type of the record",
									},
									"record_class": {
										Type: schema.TypeString, Optional: true, Description: "Class",
									},
									"record_ttl": {
										Type: schema.TypeInt, Optional: true, Description: "TTL",
									},
									"record_data": {
										Type: schema.TypeString, Optional: true, Description: "Record Data",
									},
								},
							},
						},
						"all_cached_fqdn": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cached_fqdn_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"record_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"debug_mode": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"zone_transfer": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"zone_transfer_status_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"zone_name": {
													Type: schema.TypeString, Optional: true, Description: "Zone Name",
												},
												"sflow_source_id": {
													Type: schema.TypeString, Optional: true, Description: "Sflow Source ID",
												},
												"last_update": {
													Type: schema.TypeString, Optional: true, Description: "Last Update Status",
												},
												"last_complete_update": {
													Type: schema.TypeString, Optional: true, Description: "Last Complete Update",
												},
												"last_complete_serial": {
													Type: schema.TypeString, Optional: true, Description: "Last Complete Serial",
												},
												"estimated_next_update": {
													Type: schema.TypeString, Optional: true, Description: "Estimated Next Update",
												},
											},
										},
									},
									"zone_name": {
										Type: schema.TypeString, Optional: true, Description: "Zone Name",
									},
									"sflow_source_id": {
										Type: schema.TypeString, Optional: true, Description: "Sflow Source ID",
									},
									"local_ip": {
										Type: schema.TypeString, Optional: true, Description: "Local IP",
									},
									"remote_ip": {
										Type: schema.TypeString, Optional: true, Description: "Remote IP",
									},
									"estimated_next_update": {
										Type: schema.TypeString, Optional: true, Description: "Estimated Next Update",
									},
									"zone_transfer_history_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"update_status": {
													Type: schema.TypeString, Optional: true, Description: "Update Status",
												},
												"zone_transfer_result": {
													Type: schema.TypeString, Optional: true, Description: "Zone Transfer Result",
												},
												"zone_transfer_begin_time": {
													Type: schema.TypeString, Optional: true, Description: "Zone Transfer Begin Time",
												},
												"zone_transfer_end_time": {
													Type: schema.TypeString, Optional: true, Description: "Zone Transfer End Time",
												},
												"tcp_connection_begin_time": {
													Type: schema.TypeString, Optional: true, Description: "TCP connection Begin Time",
												},
												"tcp_connection_end_time": {
													Type: schema.TypeString, Optional: true, Description: "TCP Connection End Time",
												},
												"serial_number": {
													Type: schema.TypeString, Optional: true, Description: "Serial Number",
												},
												"dns_message_processed": {
													Type: schema.TypeInt, Optional: true, Description: "DNS Message Processed",
												},
												"records_processed": {
													Type: schema.TypeInt, Optional: true, Description: "Records Processed",
												},
												"dns_message_pending_processed": {
													Type: schema.TypeInt, Optional: true, Description: "DNS message pending Processed",
												},
												"total_failure": {
													Type: schema.TypeString, Optional: true, Description: "Total Failure",
												},
											},
										},
									},
									"zone_transfer_statistics": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"stats_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"stats_count": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
									"zts_sflow_source_id": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"zone": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"statistics": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"zt_statistics": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"debug_mode": {
										Type: schema.TypeInt, Optional: true, Description: "",
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

func resourceDdosDnsCacheOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDnsCacheOperOper := setObjectDdosDnsCacheOperOper(res)
		d.Set("oper", DdosDnsCacheOperOper)
		DdosDnsCacheOperZoneTransfer := setObjectDdosDnsCacheOperZoneTransfer(res)
		d.Set("zone_transfer", DdosDnsCacheOperZoneTransfer)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDnsCacheOperOper(ret edpt.DataDdosDnsCacheOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"domain_entries":                  setSliceDdosDnsCacheOperOperDomainEntries(ret.DtDdosDnsCacheOper.Oper.DomainEntries),
			"response_status":                 ret.DtDdosDnsCacheOper.Oper.ResponseStatus,
			"response_flag":                   ret.DtDdosDnsCacheOper.Oper.ResponseFlag,
			"answer_section_record_count":     ret.DtDdosDnsCacheOper.Oper.AnswerSectionRecordCount,
			"answer_section_size":             ret.DtDdosDnsCacheOper.Oper.AnswerSectionSize,
			"authority_section_record_count":  ret.DtDdosDnsCacheOper.Oper.AuthoritySectionRecordCount,
			"authority_section_size":          ret.DtDdosDnsCacheOper.Oper.AuthoritySectionSize,
			"additional_section_record_count": ret.DtDdosDnsCacheOper.Oper.AdditionalSectionRecordCount,
			"additional_section_size":         ret.DtDdosDnsCacheOper.Oper.AdditionalSectionSize,
			"answer_section":                  setSliceDdosDnsCacheOperOperAnswerSection(ret.DtDdosDnsCacheOper.Oper.AnswerSection),
			"authoritative_section":           setSliceDdosDnsCacheOperOperAuthoritativeSection(ret.DtDdosDnsCacheOper.Oper.AuthoritativeSection),
			"additional_section":              setSliceDdosDnsCacheOperOperAdditionalSection(ret.DtDdosDnsCacheOper.Oper.AdditionalSection),
			"all_cached_fqdn":                 ret.DtDdosDnsCacheOper.Oper.AllCachedFqdn,
			"cached_fqdn_name":                ret.DtDdosDnsCacheOper.Oper.CachedFqdnName,
			"record_type":                     ret.DtDdosDnsCacheOper.Oper.RecordType,
			"debug_mode":                      ret.DtDdosDnsCacheOper.Oper.DebugMode,
		},
	}
}

func setSliceDdosDnsCacheOperOperDomainEntries(d []edpt.DdosDnsCacheOperOperDomainEntries) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["fqdn_name"] = item.FqdnName
		in["fqdn_manual_override_action"] = item.FqdnManualOverrideAction
		in["wild_card_node"] = item.WildCardNode
		in["delegation_node"] = item.DelegationNode
		in["empty_non_terminal_node"] = item.EmptyNonTerminalNode
		in["record_types"] = item.RecordTypes
		result = append(result, in)
	}
	return result
}

func setSliceDdosDnsCacheOperOperAnswerSection(d []edpt.DdosDnsCacheOperOperAnswerSection) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["record_domain_name"] = item.RecordDomainName
		in["record_type"] = item.RecordType
		in["record_class"] = item.RecordClass
		in["record_ttl"] = item.RecordTtl
		in["record_data"] = item.RecordData
		result = append(result, in)
	}
	return result
}

func setSliceDdosDnsCacheOperOperAuthoritativeSection(d []edpt.DdosDnsCacheOperOperAuthoritativeSection) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["record_domain_name"] = item.RecordDomainName
		in["record_type"] = item.RecordType
		in["record_class"] = item.RecordClass
		in["record_ttl"] = item.RecordTtl
		in["record_data"] = item.RecordData
		result = append(result, in)
	}
	return result
}

func setSliceDdosDnsCacheOperOperAdditionalSection(d []edpt.DdosDnsCacheOperOperAdditionalSection) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["record_domain_name"] = item.RecordDomainName
		in["record_type"] = item.RecordType
		in["record_class"] = item.RecordClass
		in["record_ttl"] = item.RecordTtl
		in["record_data"] = item.RecordData
		result = append(result, in)
	}
	return result
}

func setObjectDdosDnsCacheOperZoneTransfer(ret edpt.DataDdosDnsCacheOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDnsCacheOperZoneTransferOper(ret.DtDdosDnsCacheOper.ZoneTransfer.Oper),
		},
	}
}

func setObjectDdosDnsCacheOperZoneTransferOper(d edpt.DdosDnsCacheOperZoneTransferOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["zone_transfer_status_list"] = setSliceDdosDnsCacheOperZoneTransferOperZoneTransferStatusList(d.ZoneTransferStatusList)

	in["zone_name"] = d.ZoneName

	in["sflow_source_id"] = d.SflowSourceId

	in["local_ip"] = d.LocalIp

	in["remote_ip"] = d.RemoteIp

	in["estimated_next_update"] = d.EstimatedNextUpdate
	in["zone_transfer_history_list"] = setSliceDdosDnsCacheOperZoneTransferOperZoneTransferHistoryList(d.ZoneTransferHistoryList)
	in["zone_transfer_statistics"] = setSliceDdosDnsCacheOperZoneTransferOperZoneTransferStatistics(d.ZoneTransferStatistics)

	in["zts_sflow_source_id"] = d.ZtsSflowSourceId

	in["status"] = d.Status

	in["zone"] = d.Zone

	in["statistics"] = d.Statistics

	in["zt_statistics"] = d.ZtStatistics

	in["debug_mode"] = d.DebugMode
	result = append(result, in)
	return result
}

func setSliceDdosDnsCacheOperZoneTransferOperZoneTransferStatusList(d []edpt.DdosDnsCacheOperZoneTransferOperZoneTransferStatusList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["zone_name"] = item.ZoneName
		in["sflow_source_id"] = item.SflowSourceId
		in["last_update"] = item.LastUpdate
		in["last_complete_update"] = item.LastCompleteUpdate
		in["last_complete_serial"] = item.LastCompleteSerial
		in["estimated_next_update"] = item.EstimatedNextUpdate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDnsCacheOperZoneTransferOperZoneTransferHistoryList(d []edpt.DdosDnsCacheOperZoneTransferOperZoneTransferHistoryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["update_status"] = item.UpdateStatus
		in["zone_transfer_result"] = item.ZoneTransferResult
		in["zone_transfer_begin_time"] = item.ZoneTransferBeginTime
		in["zone_transfer_end_time"] = item.ZoneTransferEndTime
		in["tcp_connection_begin_time"] = item.TcpConnectionBeginTime
		in["tcp_connection_end_time"] = item.TcpConnectionEndTime
		in["serial_number"] = item.SerialNumber
		in["dns_message_processed"] = item.DnsMessageProcessed
		in["records_processed"] = item.RecordsProcessed
		in["dns_message_pending_processed"] = item.DnsMessagePendingProcessed
		in["total_failure"] = item.TotalFailure
		result = append(result, in)
	}
	return result
}

func setSliceDdosDnsCacheOperZoneTransferOperZoneTransferStatistics(d []edpt.DdosDnsCacheOperZoneTransferOperZoneTransferStatistics) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["stats_name"] = item.StatsName
		in["stats_count"] = item.StatsCount
		result = append(result, in)
	}
	return result
}

func getObjectDdosDnsCacheOperOper(d []interface{}) edpt.DdosDnsCacheOperOper {

	count1 := len(d)
	var ret edpt.DdosDnsCacheOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DomainEntries = getSliceDdosDnsCacheOperOperDomainEntries(in["domain_entries"].([]interface{}))
		ret.ResponseStatus = in["response_status"].(string)
		ret.ResponseFlag = in["response_flag"].(string)
		ret.AnswerSectionRecordCount = in["answer_section_record_count"].(int)
		ret.AnswerSectionSize = in["answer_section_size"].(int)
		ret.AuthoritySectionRecordCount = in["authority_section_record_count"].(int)
		ret.AuthoritySectionSize = in["authority_section_size"].(int)
		ret.AdditionalSectionRecordCount = in["additional_section_record_count"].(int)
		ret.AdditionalSectionSize = in["additional_section_size"].(int)
		ret.AnswerSection = getSliceDdosDnsCacheOperOperAnswerSection(in["answer_section"].([]interface{}))
		ret.AuthoritativeSection = getSliceDdosDnsCacheOperOperAuthoritativeSection(in["authoritative_section"].([]interface{}))
		ret.AdditionalSection = getSliceDdosDnsCacheOperOperAdditionalSection(in["additional_section"].([]interface{}))
		ret.AllCachedFqdn = in["all_cached_fqdn"].(int)
		ret.CachedFqdnName = in["cached_fqdn_name"].(string)
		ret.RecordType = in["record_type"].(string)
		ret.DebugMode = in["debug_mode"].(int)
	}
	return ret
}

func getSliceDdosDnsCacheOperOperDomainEntries(d []interface{}) []edpt.DdosDnsCacheOperOperDomainEntries {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheOperOperDomainEntries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheOperOperDomainEntries
		oi.FqdnName = in["fqdn_name"].(string)
		oi.FqdnManualOverrideAction = in["fqdn_manual_override_action"].(string)
		oi.WildCardNode = in["wild_card_node"].(string)
		oi.DelegationNode = in["delegation_node"].(string)
		oi.EmptyNonTerminalNode = in["empty_non_terminal_node"].(string)
		oi.RecordTypes = in["record_types"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDnsCacheOperOperAnswerSection(d []interface{}) []edpt.DdosDnsCacheOperOperAnswerSection {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheOperOperAnswerSection, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheOperOperAnswerSection
		oi.RecordDomainName = in["record_domain_name"].(string)
		oi.RecordType = in["record_type"].(string)
		oi.RecordClass = in["record_class"].(string)
		oi.RecordTtl = in["record_ttl"].(int)
		oi.RecordData = in["record_data"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDnsCacheOperOperAuthoritativeSection(d []interface{}) []edpt.DdosDnsCacheOperOperAuthoritativeSection {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheOperOperAuthoritativeSection, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheOperOperAuthoritativeSection
		oi.RecordDomainName = in["record_domain_name"].(string)
		oi.RecordType = in["record_type"].(string)
		oi.RecordClass = in["record_class"].(string)
		oi.RecordTtl = in["record_ttl"].(int)
		oi.RecordData = in["record_data"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDnsCacheOperOperAdditionalSection(d []interface{}) []edpt.DdosDnsCacheOperOperAdditionalSection {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheOperOperAdditionalSection, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheOperOperAdditionalSection
		oi.RecordDomainName = in["record_domain_name"].(string)
		oi.RecordType = in["record_type"].(string)
		oi.RecordClass = in["record_class"].(string)
		oi.RecordTtl = in["record_ttl"].(int)
		oi.RecordData = in["record_data"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDnsCacheOperZoneTransfer(d []interface{}) edpt.DdosDnsCacheOperZoneTransfer {

	count1 := len(d)
	var ret edpt.DdosDnsCacheOperZoneTransfer
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDnsCacheOperZoneTransferOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDnsCacheOperZoneTransferOper(d []interface{}) edpt.DdosDnsCacheOperZoneTransferOper {

	count1 := len(d)
	var ret edpt.DdosDnsCacheOperZoneTransferOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneTransferStatusList = getSliceDdosDnsCacheOperZoneTransferOperZoneTransferStatusList(in["zone_transfer_status_list"].([]interface{}))
		ret.ZoneName = in["zone_name"].(string)
		ret.SflowSourceId = in["sflow_source_id"].(string)
		ret.LocalIp = in["local_ip"].(string)
		ret.RemoteIp = in["remote_ip"].(string)
		ret.EstimatedNextUpdate = in["estimated_next_update"].(string)
		ret.ZoneTransferHistoryList = getSliceDdosDnsCacheOperZoneTransferOperZoneTransferHistoryList(in["zone_transfer_history_list"].([]interface{}))
		ret.ZoneTransferStatistics = getSliceDdosDnsCacheOperZoneTransferOperZoneTransferStatistics(in["zone_transfer_statistics"].([]interface{}))
		ret.ZtsSflowSourceId = in["zts_sflow_source_id"].(string)
		ret.Status = in["status"].(string)
		ret.Zone = in["zone"].(string)
		ret.Statistics = in["statistics"].(int)
		ret.ZtStatistics = in["zt_statistics"].(int)
		ret.DebugMode = in["debug_mode"].(int)
	}
	return ret
}

func getSliceDdosDnsCacheOperZoneTransferOperZoneTransferStatusList(d []interface{}) []edpt.DdosDnsCacheOperZoneTransferOperZoneTransferStatusList {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheOperZoneTransferOperZoneTransferStatusList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheOperZoneTransferOperZoneTransferStatusList
		oi.ZoneName = in["zone_name"].(string)
		oi.SflowSourceId = in["sflow_source_id"].(string)
		oi.LastUpdate = in["last_update"].(string)
		oi.LastCompleteUpdate = in["last_complete_update"].(string)
		oi.LastCompleteSerial = in["last_complete_serial"].(string)
		oi.EstimatedNextUpdate = in["estimated_next_update"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDnsCacheOperZoneTransferOperZoneTransferHistoryList(d []interface{}) []edpt.DdosDnsCacheOperZoneTransferOperZoneTransferHistoryList {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheOperZoneTransferOperZoneTransferHistoryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheOperZoneTransferOperZoneTransferHistoryList
		oi.UpdateStatus = in["update_status"].(string)
		oi.ZoneTransferResult = in["zone_transfer_result"].(string)
		oi.ZoneTransferBeginTime = in["zone_transfer_begin_time"].(string)
		oi.ZoneTransferEndTime = in["zone_transfer_end_time"].(string)
		oi.TcpConnectionBeginTime = in["tcp_connection_begin_time"].(string)
		oi.TcpConnectionEndTime = in["tcp_connection_end_time"].(string)
		oi.SerialNumber = in["serial_number"].(string)
		oi.DnsMessageProcessed = in["dns_message_processed"].(int)
		oi.RecordsProcessed = in["records_processed"].(int)
		oi.DnsMessagePendingProcessed = in["dns_message_pending_processed"].(int)
		oi.TotalFailure = in["total_failure"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDnsCacheOperZoneTransferOperZoneTransferStatistics(d []interface{}) []edpt.DdosDnsCacheOperZoneTransferOperZoneTransferStatistics {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheOperZoneTransferOperZoneTransferStatistics, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheOperZoneTransferOperZoneTransferStatistics
		oi.StatsName = in["stats_name"].(string)
		oi.StatsCount = in["stats_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDnsCacheOper(d *schema.ResourceData) edpt.DdosDnsCacheOper {
	var ret edpt.DdosDnsCacheOper

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectDdosDnsCacheOperOper(d.Get("oper").([]interface{}))

	ret.ZoneTransfer = getObjectDdosDnsCacheOperZoneTransfer(d.Get("zone_transfer").([]interface{}))
	return ret
}
