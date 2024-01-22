package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneDetectionOutboundDetectionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_detection_outbound_detection_oper`: Operational Status for the object outbound-detection\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneDetectionOutboundDetectionOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"discovery_timestamp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"location_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"location_name": {
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
												"rate": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"maximum": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"minimum": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"non_zero_minimum": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"average": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"adaptive_threshold": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"data_source": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"anomaly": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"anomaly_timestamp": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"initial_learning": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"active_time": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"topk_source_subnet": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"entry_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"location_type": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"location_name": {
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
															"source_subnets": {
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
											},
										},
									},
								},
							},
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

func resourceDdosDstZoneDetectionOutboundDetectionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionOutboundDetectionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionOutboundDetectionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneDetectionOutboundDetectionOperOper := setObjectDdosDstZoneDetectionOutboundDetectionOperOper(res)
		d.Set("oper", DdosDstZoneDetectionOutboundDetectionOperOper)
		DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnet := setObjectDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnet(res)
		d.Set("topk_source_subnet", DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnet)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneDetectionOutboundDetectionOperOper(ret edpt.DataDdosDstZoneDetectionOutboundDetectionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"discovery_timestamp": ret.DtDdosDstZoneDetectionOutboundDetectionOper.Oper.DiscoveryTimestamp,
			"entry_list":          setSliceDdosDstZoneDetectionOutboundDetectionOperOperEntryList(ret.DtDdosDstZoneDetectionOutboundDetectionOper.Oper.EntryList),
		},
	}
}

func setSliceDdosDstZoneDetectionOutboundDetectionOperOperEntryList(d []edpt.DdosDstZoneDetectionOutboundDetectionOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["location_type"] = item.LocationType
		in["location_name"] = item.LocationName
		in["indicators"] = setSliceDdosDstZoneDetectionOutboundDetectionOperOperEntryListIndicators(item.Indicators)
		in["data_source"] = item.DataSource
		in["anomaly"] = item.Anomaly
		in["anomaly_timestamp"] = item.AnomalyTimestamp
		in["initial_learning"] = item.InitialLearning
		in["active_time"] = item.ActiveTime
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneDetectionOutboundDetectionOperOperEntryListIndicators(d []edpt.DdosDstZoneDetectionOutboundDetectionOperOperEntryListIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["rate"] = item.Rate
		in["maximum"] = item.Maximum
		in["minimum"] = item.Minimum
		in["non_zero_minimum"] = item.NonZeroMinimum
		in["average"] = item.Average
		in["adaptive_threshold"] = item.AdaptiveThreshold
		result = append(result, in)
	}
	return result
}

func setObjectDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnet(ret edpt.DataDdosDstZoneDetectionOutboundDetectionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOper(ret.DtDdosDstZoneDetectionOutboundDetectionOper.TopkSourceSubnet.Oper),
		},
	}
}

func setObjectDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOper(d edpt.DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["entry_list"] = setSliceDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryList(d.EntryList)
	result = append(result, in)
	return result
}

func setSliceDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryList(d []edpt.DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["location_type"] = item.LocationType
		in["location_name"] = item.LocationName
		in["indicators"] = setSliceDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicators(d []edpt.DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["source_subnets"] = setSliceDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicatorsSourceSubnets(item.SourceSubnets)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicatorsSourceSubnets(d []edpt.DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicatorsSourceSubnets) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZoneDetectionOutboundDetectionOperOper(d []interface{}) edpt.DdosDstZoneDetectionOutboundDetectionOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneDetectionOutboundDetectionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DiscoveryTimestamp = in["discovery_timestamp"].(string)
		ret.EntryList = getSliceDdosDstZoneDetectionOutboundDetectionOperOperEntryList(in["entry_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneDetectionOutboundDetectionOperOperEntryList(d []interface{}) []edpt.DdosDstZoneDetectionOutboundDetectionOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionOutboundDetectionOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionOutboundDetectionOperOperEntryList
		oi.LocationType = in["location_type"].(string)
		oi.LocationName = in["location_name"].(string)
		oi.Indicators = getSliceDdosDstZoneDetectionOutboundDetectionOperOperEntryListIndicators(in["indicators"].([]interface{}))
		oi.DataSource = in["data_source"].(string)
		oi.Anomaly = in["anomaly"].(string)
		oi.AnomalyTimestamp = in["anomaly_timestamp"].(string)
		oi.InitialLearning = in["initial_learning"].(string)
		oi.ActiveTime = in["active_time"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneDetectionOutboundDetectionOperOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneDetectionOutboundDetectionOperOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionOutboundDetectionOperOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionOutboundDetectionOperOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.Maximum = in["maximum"].(string)
		oi.Minimum = in["minimum"].(string)
		oi.NonZeroMinimum = in["non_zero_minimum"].(string)
		oi.Average = in["average"].(string)
		oi.AdaptiveThreshold = in["adaptive_threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnet(d []interface{}) edpt.DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnet {

	count1 := len(d)
	var ret edpt.DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnet
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOper(d []interface{}) edpt.DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EntryList = getSliceDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryList(in["entry_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryList(d []interface{}) []edpt.DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryList
		oi.LocationType = in["location_type"].(string)
		oi.LocationName = in["location_name"].(string)
		oi.Indicators = getSliceDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.SourceSubnets = getSliceDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicatorsSourceSubnets(in["source_subnets"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicatorsSourceSubnets(d []interface{}) []edpt.DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicatorsSourceSubnets {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicatorsSourceSubnets, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnetOperEntryListIndicatorsSourceSubnets
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneDetectionOutboundDetectionOper(d *schema.ResourceData) edpt.DdosDstZoneDetectionOutboundDetectionOper {
	var ret edpt.DdosDstZoneDetectionOutboundDetectionOper

	ret.Oper = getObjectDdosDstZoneDetectionOutboundDetectionOperOper(d.Get("oper").([]interface{}))

	ret.TopkSourceSubnet = getObjectDdosDstZoneDetectionOutboundDetectionOperTopkSourceSubnet(d.Get("topk_source_subnet").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)
	return ret
}
