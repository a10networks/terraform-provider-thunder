package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_network_object_oper`: Operational Status for the object network-object\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosNetworkObjectOperRead,

		Schema: map[string]*schema.Schema{
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"children_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"port_range_start": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"port_range_end": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"service_protocol": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"agent_group_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
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
												"value": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"current": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"threshold": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"max": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
												"is_anomaly": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
									"is_anomaly": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"is_learning_done": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"is_histogram_learning_done": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"operational_mode": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"histogram_mode": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"display_filter": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"es_timestamp": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"de_es_timestamp": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"details": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"trusted_details": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_details": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"victim_list": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"discovered_list": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sport_list": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"subnet_ip_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"subnet_ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv4": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"discovered_ip_list": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"anomaly_ip_list": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sport": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"port_start": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"port_end": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"protocol": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"single_layer_discovered_list": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"agent_group_details": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"topk_destinations": {
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
												"destinations": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"address": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"rate": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
											},
										},
									},
									"topk_type": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reset_time": {
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

func resourceDdosNetworkObjectOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosNetworkObjectOperOper := setObjectDdosNetworkObjectOperOper(res)
		d.Set("oper", DdosNetworkObjectOperOper)
		DdosNetworkObjectOperTopkDestinations := setObjectDdosNetworkObjectOperTopkDestinations(res)
		d.Set("topk_destinations", DdosNetworkObjectOperTopkDestinations)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosNetworkObjectOperOper(ret edpt.DataDdosNetworkObjectOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entry_list":                   setSliceDdosNetworkObjectOperOperEntryList(ret.DtDdosNetworkObjectOper.Oper.EntryList),
			"entry_count":                  ret.DtDdosNetworkObjectOper.Oper.EntryCount,
			"details":                      ret.DtDdosNetworkObjectOper.Oper.Details,
			"trusted_details":              ret.DtDdosNetworkObjectOper.Oper.TrustedDetails,
			"total_details":                ret.DtDdosNetworkObjectOper.Oper.TotalDetails,
			"victim_list":                  ret.DtDdosNetworkObjectOper.Oper.VictimList,
			"discovered_list":              ret.DtDdosNetworkObjectOper.Oper.DiscoveredList,
			"sport_list":                   ret.DtDdosNetworkObjectOper.Oper.SportList,
			"subnet_ip_addr":               ret.DtDdosNetworkObjectOper.Oper.SubnetIpAddr,
			"subnet_ipv6_addr":             ret.DtDdosNetworkObjectOper.Oper.SubnetIpv6Addr,
			"ipv4":                         ret.DtDdosNetworkObjectOper.Oper.Ipv4,
			"discovered_ip_list":           ret.DtDdosNetworkObjectOper.Oper.DiscoveredIpList,
			"anomaly_ip_list":              ret.DtDdosNetworkObjectOper.Oper.AnomalyIpList,
			"sport":                        ret.DtDdosNetworkObjectOper.Oper.Sport,
			"port_start":                   ret.DtDdosNetworkObjectOper.Oper.PortStart,
			"port_end":                     ret.DtDdosNetworkObjectOper.Oper.PortEnd,
			"protocol":                     ret.DtDdosNetworkObjectOper.Oper.Protocol,
			"single_layer_discovered_list": ret.DtDdosNetworkObjectOper.Oper.SingleLayerDiscoveredList,
			"agent_group_details":          ret.DtDdosNetworkObjectOper.Oper.AgentGroupDetails,
		},
	}
}

func setSliceDdosNetworkObjectOperOperEntryList(d []edpt.DdosNetworkObjectOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["children_count"] = item.ChildrenCount
		in["port_range_start"] = item.PortRangeStart
		in["port_range_end"] = item.PortRangeEnd
		in["port"] = item.Port
		in["service_protocol"] = item.ServiceProtocol
		in["agent_group_name"] = item.AgentGroupName
		in["indicators"] = setSliceDdosNetworkObjectOperOperEntryListIndicators(item.Indicators)
		in["is_anomaly"] = item.IsAnomaly
		in["is_learning_done"] = item.IsLearningDone
		in["is_histogram_learning_done"] = item.IsHistogramLearningDone
		in["operational_mode"] = item.OperationalMode
		in["histogram_mode"] = item.HistogramMode
		in["display_filter"] = item.DisplayFilter
		in["es_timestamp"] = item.EsTimestamp
		in["de_es_timestamp"] = item.DeEsTimestamp
		result = append(result, in)
	}
	return result
}

func setSliceDdosNetworkObjectOperOperEntryListIndicators(d []edpt.DdosNetworkObjectOperOperEntryListIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["value"] = setSliceDdosNetworkObjectOperOperEntryListIndicatorsValue(item.Value)
		in["is_anomaly"] = item.IsAnomaly
		result = append(result, in)
	}
	return result
}

func setSliceDdosNetworkObjectOperOperEntryListIndicatorsValue(d []edpt.DdosNetworkObjectOperOperEntryListIndicatorsValue) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["current"] = item.Current
		in["threshold"] = item.Threshold
		in["max"] = item.Max
		result = append(result, in)
	}
	return result
}

func setObjectDdosNetworkObjectOperTopkDestinations(ret edpt.DataDdosNetworkObjectOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosNetworkObjectOperTopkDestinationsOper(ret.DtDdosNetworkObjectOper.TopkDestinations.Oper),
		},
	}
}

func setObjectDdosNetworkObjectOperTopkDestinationsOper(d edpt.DdosNetworkObjectOperTopkDestinationsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["indicators"] = setSliceDdosNetworkObjectOperTopkDestinationsOperIndicators(d.Indicators)

	in["topk_type"] = d.TopkType

	in["reset_time"] = d.ResetTime
	result = append(result, in)
	return result
}

func setSliceDdosNetworkObjectOperTopkDestinationsOperIndicators(d []edpt.DdosNetworkObjectOperTopkDestinationsOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["destinations"] = setSliceDdosNetworkObjectOperTopkDestinationsOperIndicatorsDestinations(item.Destinations)
		result = append(result, in)
	}
	return result
}

func setSliceDdosNetworkObjectOperTopkDestinationsOperIndicatorsDestinations(d []edpt.DdosNetworkObjectOperTopkDestinationsOperIndicatorsDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func getObjectDdosNetworkObjectOperOper(d []interface{}) edpt.DdosNetworkObjectOperOper {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EntryList = getSliceDdosNetworkObjectOperOperEntryList(in["entry_list"].([]interface{}))
		ret.EntryCount = in["entry_count"].(int)
		ret.Details = in["details"].(int)
		ret.TrustedDetails = in["trusted_details"].(int)
		ret.TotalDetails = in["total_details"].(int)
		ret.VictimList = in["victim_list"].(int)
		ret.DiscoveredList = in["discovered_list"].(int)
		ret.SportList = in["sport_list"].(int)
		ret.SubnetIpAddr = in["subnet_ip_addr"].(string)
		ret.SubnetIpv6Addr = in["subnet_ipv6_addr"].(string)
		ret.Ipv4 = in["ipv4"].(string)
		ret.DiscoveredIpList = in["discovered_ip_list"].(int)
		ret.AnomalyIpList = in["anomaly_ip_list"].(int)
		ret.Sport = in["sport"].(int)
		ret.PortStart = in["port_start"].(int)
		ret.PortEnd = in["port_end"].(int)
		ret.Protocol = in["protocol"].(int)
		ret.SingleLayerDiscoveredList = in["single_layer_discovered_list"].(int)
		ret.AgentGroupDetails = in["agent_group_details"].(int)
	}
	return ret
}

func getSliceDdosNetworkObjectOperOperEntryList(d []interface{}) []edpt.DdosNetworkObjectOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectOperOperEntryList
		oi.Address = in["address"].(string)
		oi.ChildrenCount = in["children_count"].(int)
		oi.PortRangeStart = in["port_range_start"].(int)
		oi.PortRangeEnd = in["port_range_end"].(int)
		oi.Port = in["port"].(int)
		oi.ServiceProtocol = in["service_protocol"].(string)
		oi.AgentGroupName = in["agent_group_name"].(string)
		oi.Indicators = getSliceDdosNetworkObjectOperOperEntryListIndicators(in["indicators"].([]interface{}))
		oi.IsAnomaly = in["is_anomaly"].(int)
		oi.IsLearningDone = in["is_learning_done"].(int)
		oi.IsHistogramLearningDone = in["is_histogram_learning_done"].(int)
		oi.OperationalMode = in["operational_mode"].(int)
		oi.HistogramMode = in["histogram_mode"].(int)
		oi.DisplayFilter = in["display_filter"].(string)
		oi.EsTimestamp = in["es_timestamp"].(string)
		oi.DeEsTimestamp = in["de_es_timestamp"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosNetworkObjectOperOperEntryListIndicators(d []interface{}) []edpt.DdosNetworkObjectOperOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectOperOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectOperOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Value = getSliceDdosNetworkObjectOperOperEntryListIndicatorsValue(in["value"].([]interface{}))
		oi.IsAnomaly = in["is_anomaly"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosNetworkObjectOperOperEntryListIndicatorsValue(d []interface{}) []edpt.DdosNetworkObjectOperOperEntryListIndicatorsValue {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectOperOperEntryListIndicatorsValue, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectOperOperEntryListIndicatorsValue
		oi.Current = in["current"].(string)
		oi.Threshold = in["threshold"].(string)
		oi.Max = in["max"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectOperTopkDestinations(d []interface{}) edpt.DdosNetworkObjectOperTopkDestinations {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectOperTopkDestinations
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosNetworkObjectOperTopkDestinationsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosNetworkObjectOperTopkDestinationsOper(d []interface{}) edpt.DdosNetworkObjectOperTopkDestinationsOper {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectOperTopkDestinationsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosNetworkObjectOperTopkDestinationsOperIndicators(in["indicators"].([]interface{}))
		ret.TopkType = in["topk_type"].(int)
		ret.ResetTime = in["reset_time"].(int)
	}
	return ret
}

func getSliceDdosNetworkObjectOperTopkDestinationsOperIndicators(d []interface{}) []edpt.DdosNetworkObjectOperTopkDestinationsOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectOperTopkDestinationsOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectOperTopkDestinationsOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Destinations = getSliceDdosNetworkObjectOperTopkDestinationsOperIndicatorsDestinations(in["destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosNetworkObjectOperTopkDestinationsOperIndicatorsDestinations(d []interface{}) []edpt.DdosNetworkObjectOperTopkDestinationsOperIndicatorsDestinations {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectOperTopkDestinationsOperIndicatorsDestinations, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectOperTopkDestinationsOperIndicatorsDestinations
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosNetworkObjectOper(d *schema.ResourceData) edpt.DdosNetworkObjectOper {
	var ret edpt.DdosNetworkObjectOper

	ret.ObjectName = d.Get("object_name").(string)

	ret.Oper = getObjectDdosNetworkObjectOperOper(d.Get("oper").([]interface{}))

	ret.TopkDestinations = getObjectDdosNetworkObjectOperTopkDestinations(d.Get("topk_destinations").([]interface{}))
	return ret
}
