package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDnsCacheZoneTransferOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dns_cache_zone_transfer_oper`: Operational Status for the object zone-transfer\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDnsCacheZoneTransferOperRead,

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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceDdosDnsCacheZoneTransferOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheZoneTransferOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheZoneTransferOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDnsCacheZoneTransferOperOper := setObjectDdosDnsCacheZoneTransferOperOper(res)
		d.Set("oper", DdosDnsCacheZoneTransferOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDnsCacheZoneTransferOperOper(ret edpt.DataDdosDnsCacheZoneTransferOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"zone_transfer_status_list":  setSliceDdosDnsCacheZoneTransferOperOperZoneTransferStatusList(ret.DtDdosDnsCacheZoneTransferOper.Oper.ZoneTransferStatusList),
			"zone_name":                  ret.DtDdosDnsCacheZoneTransferOper.Oper.ZoneName,
			"sflow_source_id":            ret.DtDdosDnsCacheZoneTransferOper.Oper.SflowSourceId,
			"local_ip":                   ret.DtDdosDnsCacheZoneTransferOper.Oper.LocalIp,
			"remote_ip":                  ret.DtDdosDnsCacheZoneTransferOper.Oper.RemoteIp,
			"estimated_next_update":      ret.DtDdosDnsCacheZoneTransferOper.Oper.EstimatedNextUpdate,
			"zone_transfer_history_list": setSliceDdosDnsCacheZoneTransferOperOperZoneTransferHistoryList(ret.DtDdosDnsCacheZoneTransferOper.Oper.ZoneTransferHistoryList),
			"zone_transfer_statistics":   setSliceDdosDnsCacheZoneTransferOperOperZoneTransferStatistics(ret.DtDdosDnsCacheZoneTransferOper.Oper.ZoneTransferStatistics),
			"zts_sflow_source_id":        ret.DtDdosDnsCacheZoneTransferOper.Oper.ZtsSflowSourceId,
			"status":                     ret.DtDdosDnsCacheZoneTransferOper.Oper.Status,
			"zone":                       ret.DtDdosDnsCacheZoneTransferOper.Oper.Zone,
			"statistics":                 ret.DtDdosDnsCacheZoneTransferOper.Oper.Statistics,
			"zt_statistics":              ret.DtDdosDnsCacheZoneTransferOper.Oper.ZtStatistics,
			"debug_mode":                 ret.DtDdosDnsCacheZoneTransferOper.Oper.DebugMode,
		},
	}
}

func setSliceDdosDnsCacheZoneTransferOperOperZoneTransferStatusList(d []edpt.DdosDnsCacheZoneTransferOperOperZoneTransferStatusList) []map[string]interface{} {
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

func setSliceDdosDnsCacheZoneTransferOperOperZoneTransferHistoryList(d []edpt.DdosDnsCacheZoneTransferOperOperZoneTransferHistoryList) []map[string]interface{} {
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

func setSliceDdosDnsCacheZoneTransferOperOperZoneTransferStatistics(d []edpt.DdosDnsCacheZoneTransferOperOperZoneTransferStatistics) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["stats_name"] = item.StatsName
		in["stats_count"] = item.StatsCount
		result = append(result, in)
	}
	return result
}

func getObjectDdosDnsCacheZoneTransferOperOper(d []interface{}) edpt.DdosDnsCacheZoneTransferOperOper {

	count1 := len(d)
	var ret edpt.DdosDnsCacheZoneTransferOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneTransferStatusList = getSliceDdosDnsCacheZoneTransferOperOperZoneTransferStatusList(in["zone_transfer_status_list"].([]interface{}))
		ret.ZoneName = in["zone_name"].(string)
		ret.SflowSourceId = in["sflow_source_id"].(string)
		ret.LocalIp = in["local_ip"].(string)
		ret.RemoteIp = in["remote_ip"].(string)
		ret.EstimatedNextUpdate = in["estimated_next_update"].(string)
		ret.ZoneTransferHistoryList = getSliceDdosDnsCacheZoneTransferOperOperZoneTransferHistoryList(in["zone_transfer_history_list"].([]interface{}))
		ret.ZoneTransferStatistics = getSliceDdosDnsCacheZoneTransferOperOperZoneTransferStatistics(in["zone_transfer_statistics"].([]interface{}))
		ret.ZtsSflowSourceId = in["zts_sflow_source_id"].(string)
		ret.Status = in["status"].(string)
		ret.Zone = in["zone"].(string)
		ret.Statistics = in["statistics"].(int)
		ret.ZtStatistics = in["zt_statistics"].(int)
		ret.DebugMode = in["debug_mode"].(int)
	}
	return ret
}

func getSliceDdosDnsCacheZoneTransferOperOperZoneTransferStatusList(d []interface{}) []edpt.DdosDnsCacheZoneTransferOperOperZoneTransferStatusList {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheZoneTransferOperOperZoneTransferStatusList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheZoneTransferOperOperZoneTransferStatusList
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

func getSliceDdosDnsCacheZoneTransferOperOperZoneTransferHistoryList(d []interface{}) []edpt.DdosDnsCacheZoneTransferOperOperZoneTransferHistoryList {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheZoneTransferOperOperZoneTransferHistoryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheZoneTransferOperOperZoneTransferHistoryList
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

func getSliceDdosDnsCacheZoneTransferOperOperZoneTransferStatistics(d []interface{}) []edpt.DdosDnsCacheZoneTransferOperOperZoneTransferStatistics {

	count1 := len(d)
	ret := make([]edpt.DdosDnsCacheZoneTransferOperOperZoneTransferStatistics, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDnsCacheZoneTransferOperOperZoneTransferStatistics
		oi.StatsName = in["stats_name"].(string)
		oi.StatsCount = in["stats_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDnsCacheZoneTransferOper(d *schema.ResourceData) edpt.DdosDnsCacheZoneTransferOper {
	var ret edpt.DdosDnsCacheZoneTransferOper

	ret.Oper = getObjectDdosDnsCacheZoneTransferOperOper(d.Get("oper").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
