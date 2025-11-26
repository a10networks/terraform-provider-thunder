package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneOutboundPolicyOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_outbound_policy_oper`: Operational Status for the object outbound-policy\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneOutboundPolicyOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"no_class_list_match": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"policy_class_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_list_name": {
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
									"current_frag_packet_rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_frag_packet_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"frag_packet_rate_limit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"age_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"lockup_time": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"packet_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"packet_dropped": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"packet_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"kbit_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"kbit_rate_exceed_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connections_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connection_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"frag_packet_rate": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"geo_tracking_statistics": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"packet_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"packet_dropped": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"packet_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"kbit_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"kbit_rate_exceed_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connections_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connection_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"frag_packet_rate": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tracking_entry_learn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tracking_entry_aged": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tracking_entry_learning_thre_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"tracking_entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"geo_location_name": {
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
								},
							},
						},
						"policy_rate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"policy_statistics": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"tracking_entry_filter": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}

func resourceDdosDstZoneOutboundPolicyOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneOutboundPolicyOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneOutboundPolicyOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneOutboundPolicyOperOper := setObjectDdosDstZoneOutboundPolicyOperOper(res)
		d.Set("oper", DdosDstZoneOutboundPolicyOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneOutboundPolicyOperOper(ret edpt.DataDdosDstZoneOutboundPolicyOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"policy_name":             ret.DtDdosDstZoneOutboundPolicyOper.Oper.PolicyName,
			"no_class_list_match":     ret.DtDdosDstZoneOutboundPolicyOper.Oper.NoClassListMatch,
			"policy_class_list":       setSliceDdosDstZoneOutboundPolicyOperOperPolicyClassList(ret.DtDdosDstZoneOutboundPolicyOper.Oper.PolicyClassList),
			"geo_tracking_statistics": setObjectDdosDstZoneOutboundPolicyOperOperGeoTrackingStatistics(ret.DtDdosDstZoneOutboundPolicyOper.Oper.GeoTrackingStatistics),
			"tracking_entry_list":     setSliceDdosDstZoneOutboundPolicyOperOperTrackingEntryList(ret.DtDdosDstZoneOutboundPolicyOper.Oper.TrackingEntryList),
			"policy_rate":             ret.DtDdosDstZoneOutboundPolicyOper.Oper.PolicyRate,
			"policy_statistics":       ret.DtDdosDstZoneOutboundPolicyOper.Oper.PolicyStatistics,
			"tracking_entry_filter":   ret.DtDdosDstZoneOutboundPolicyOper.Oper.TrackingEntryFilter,
		},
	}
}

func setSliceDdosDstZoneOutboundPolicyOperOperPolicyClassList(d []edpt.DdosDstZoneOutboundPolicyOperOperPolicyClassList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["class_list_name"] = item.ClassListName
		in["current_packet_rate"] = item.CurrentPacketRate
		in["is_packet_rate_exceed"] = item.IsPacketRateExceed
		in["packet_rate_limit"] = item.PacketRateLimit
		in["current_kbit_rate"] = item.CurrentKbitRate
		in["is_kbit_rate_exceed"] = item.IsKbitRateExceed
		in["kbit_rate_limit"] = item.KbitRateLimit
		in["current_connections"] = item.CurrentConnections
		in["is_connections_exceed"] = item.IsConnectionsExceed
		in["connection_limit"] = item.ConnectionLimit
		in["current_connection_rate"] = item.CurrentConnectionRate
		in["is_connection_rate_exceed"] = item.IsConnectionRateExceed
		in["connection_rate_limit"] = item.ConnectionRateLimit
		in["current_frag_packet_rate"] = item.CurrentFragPacketRate
		in["is_frag_packet_rate_exceed"] = item.IsFragPacketRateExceed
		in["frag_packet_rate_limit"] = item.FragPacketRateLimit
		in["age_str"] = item.AgeStr
		in["lockup_time"] = item.LockupTime
		in["packet_received"] = item.PacketReceived
		in["packet_dropped"] = item.PacketDropped
		in["packet_rate_exceed"] = item.PacketRateExceed
		in["kbit_rate_exceed"] = item.KbitRateExceed
		in["kbit_rate_exceed_count"] = item.KbitRateExceedCount
		in["connections_exceed"] = item.ConnectionsExceed
		in["connection_rate_exceed"] = item.ConnectionRateExceed
		in["frag_packet_rate"] = item.FragPacketRate
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneOutboundPolicyOperOperGeoTrackingStatistics(d edpt.DdosDstZoneOutboundPolicyOperOperGeoTrackingStatistics) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["packet_received"] = d.PacketReceived

	in["packet_dropped"] = d.PacketDropped

	in["packet_rate_exceed"] = d.PacketRateExceed

	in["kbit_rate_exceed"] = d.KbitRateExceed

	in["kbit_rate_exceed_count"] = d.KbitRateExceedCount

	in["connections_exceed"] = d.ConnectionsExceed

	in["connection_rate_exceed"] = d.ConnectionRateExceed

	in["frag_packet_rate"] = d.FragPacketRate

	in["tracking_entry_learn"] = d.TrackingEntryLearn

	in["tracking_entry_aged"] = d.TrackingEntryAged

	in["tracking_entry_learning_thre_exceed"] = d.TrackingEntryLearningThreExceed
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneOutboundPolicyOperOperTrackingEntryList(d []edpt.DdosDstZoneOutboundPolicyOperOperTrackingEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["geo_location_name"] = item.GeoLocationName
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
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZoneOutboundPolicyOperOper(d []interface{}) edpt.DdosDstZoneOutboundPolicyOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneOutboundPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PolicyName = in["policy_name"].(string)
		ret.NoClassListMatch = in["no_class_list_match"].(int)
		ret.PolicyClassList = getSliceDdosDstZoneOutboundPolicyOperOperPolicyClassList(in["policy_class_list"].([]interface{}))
		ret.GeoTrackingStatistics = getObjectDdosDstZoneOutboundPolicyOperOperGeoTrackingStatistics(in["geo_tracking_statistics"].([]interface{}))
		ret.TrackingEntryList = getSliceDdosDstZoneOutboundPolicyOperOperTrackingEntryList(in["tracking_entry_list"].([]interface{}))
		ret.PolicyRate = in["policy_rate"].(int)
		ret.PolicyStatistics = in["policy_statistics"].(int)
		ret.TrackingEntryFilter = in["tracking_entry_filter"].(int)
	}
	return ret
}

func getSliceDdosDstZoneOutboundPolicyOperOperPolicyClassList(d []interface{}) []edpt.DdosDstZoneOutboundPolicyOperOperPolicyClassList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOutboundPolicyOperOperPolicyClassList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOutboundPolicyOperOperPolicyClassList
		oi.ClassListName = in["class_list_name"].(string)
		oi.CurrentPacketRate = in["current_packet_rate"].(string)
		oi.IsPacketRateExceed = in["is_packet_rate_exceed"].(int)
		oi.PacketRateLimit = in["packet_rate_limit"].(string)
		oi.CurrentKbitRate = in["current_kbit_rate"].(string)
		oi.IsKbitRateExceed = in["is_kbit_rate_exceed"].(int)
		oi.KbitRateLimit = in["kbit_rate_limit"].(string)
		oi.CurrentConnections = in["current_connections"].(string)
		oi.IsConnectionsExceed = in["is_connections_exceed"].(int)
		oi.ConnectionLimit = in["connection_limit"].(string)
		oi.CurrentConnectionRate = in["current_connection_rate"].(string)
		oi.IsConnectionRateExceed = in["is_connection_rate_exceed"].(int)
		oi.ConnectionRateLimit = in["connection_rate_limit"].(string)
		oi.CurrentFragPacketRate = in["current_frag_packet_rate"].(string)
		oi.IsFragPacketRateExceed = in["is_frag_packet_rate_exceed"].(int)
		oi.FragPacketRateLimit = in["frag_packet_rate_limit"].(string)
		oi.AgeStr = in["age_str"].(string)
		oi.LockupTime = in["lockup_time"].(int)
		oi.PacketReceived = in["packet_received"].(int)
		oi.PacketDropped = in["packet_dropped"].(int)
		oi.PacketRateExceed = in["packet_rate_exceed"].(int)
		oi.KbitRateExceed = in["kbit_rate_exceed"].(int)
		oi.KbitRateExceedCount = in["kbit_rate_exceed_count"].(int)
		oi.ConnectionsExceed = in["connections_exceed"].(int)
		oi.ConnectionRateExceed = in["connection_rate_exceed"].(int)
		oi.FragPacketRate = in["frag_packet_rate"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneOutboundPolicyOperOperGeoTrackingStatistics(d []interface{}) edpt.DdosDstZoneOutboundPolicyOperOperGeoTrackingStatistics {

	count1 := len(d)
	var ret edpt.DdosDstZoneOutboundPolicyOperOperGeoTrackingStatistics
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PacketReceived = in["packet_received"].(int)
		ret.PacketDropped = in["packet_dropped"].(int)
		ret.PacketRateExceed = in["packet_rate_exceed"].(int)
		ret.KbitRateExceed = in["kbit_rate_exceed"].(int)
		ret.KbitRateExceedCount = in["kbit_rate_exceed_count"].(int)
		ret.ConnectionsExceed = in["connections_exceed"].(int)
		ret.ConnectionRateExceed = in["connection_rate_exceed"].(int)
		ret.FragPacketRate = in["frag_packet_rate"].(int)
		ret.TrackingEntryLearn = in["tracking_entry_learn"].(int)
		ret.TrackingEntryAged = in["tracking_entry_aged"].(int)
		ret.TrackingEntryLearningThreExceed = in["tracking_entry_learning_thre_exceed"].(int)
	}
	return ret
}

func getSliceDdosDstZoneOutboundPolicyOperOperTrackingEntryList(d []interface{}) []edpt.DdosDstZoneOutboundPolicyOperOperTrackingEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneOutboundPolicyOperOperTrackingEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneOutboundPolicyOperOperTrackingEntryList
		oi.GeoLocationName = in["geo_location_name"].(string)
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
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneOutboundPolicyOper(d *schema.ResourceData) edpt.DdosDstZoneOutboundPolicyOper {
	var ret edpt.DdosDstZoneOutboundPolicyOper

	ret.Oper = getObjectDdosDstZoneOutboundPolicyOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)
	return ret
}
