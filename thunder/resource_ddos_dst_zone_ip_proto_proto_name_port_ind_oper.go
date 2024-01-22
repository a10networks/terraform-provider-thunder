package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNamePortIndOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_ip_proto_proto_name_port_ind_oper`: Operational Status for the object port-ind\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneIpProtoProtoNamePortIndOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_address_str": {
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
												"src_maximum": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"src_minimum": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"src_non_zero_minimum": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"src_average": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"score": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"detection_data_source": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"total_score": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"current_level": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"src_level": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"escalation_timestamp": {
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
									"zone_maximum": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"zone_minimum": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"zone_non_zero_minimum": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"zone_average": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"zone_adaptive_threshold": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"src_maximum": {
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
									"score": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"detection_data_source": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"total_score": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"current_level": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"escalation_timestamp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"initial_learning": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"active_time": {
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
						"details": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sources": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
		},
	}
}

func resourceDdosDstZoneIpProtoProtoNamePortIndOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNamePortIndOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNamePortIndOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneIpProtoProtoNamePortIndOperOper := setObjectDdosDstZoneIpProtoProtoNamePortIndOperOper(res)
		d.Set("oper", DdosDstZoneIpProtoProtoNamePortIndOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneIpProtoProtoNamePortIndOperOper(ret edpt.DataDdosDstZoneIpProtoProtoNamePortIndOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"src_entry_list":        setSliceDdosDstZoneIpProtoProtoNamePortIndOperOperSrcEntryList(ret.DtDdosDstZoneIpProtoProtoNamePortIndOper.Oper.SrcEntryList),
			"indicators":            setSliceDdosDstZoneIpProtoProtoNamePortIndOperOperIndicators(ret.DtDdosDstZoneIpProtoProtoNamePortIndOper.Oper.Indicators),
			"detection_data_source": ret.DtDdosDstZoneIpProtoProtoNamePortIndOper.Oper.DetectionDataSource,
			"total_score":           ret.DtDdosDstZoneIpProtoProtoNamePortIndOper.Oper.TotalScore,
			"current_level":         ret.DtDdosDstZoneIpProtoProtoNamePortIndOper.Oper.CurrentLevel,
			"escalation_timestamp":  ret.DtDdosDstZoneIpProtoProtoNamePortIndOper.Oper.EscalationTimestamp,
			"initial_learning":      ret.DtDdosDstZoneIpProtoProtoNamePortIndOper.Oper.InitialLearning,
			"active_time":           ret.DtDdosDstZoneIpProtoProtoNamePortIndOper.Oper.ActiveTime,
			"sources_all_entries":   ret.DtDdosDstZoneIpProtoProtoNamePortIndOper.Oper.SourcesAllEntries,
			"subnet_ip_addr":        ret.DtDdosDstZoneIpProtoProtoNamePortIndOper.Oper.SubnetIpAddr,
			"subnet_ipv6_addr":      ret.DtDdosDstZoneIpProtoProtoNamePortIndOper.Oper.SubnetIpv6Addr,
			"ipv6":                  ret.DtDdosDstZoneIpProtoProtoNamePortIndOper.Oper.Ipv6,
			"details":               ret.DtDdosDstZoneIpProtoProtoNamePortIndOper.Oper.Details,
			"sources":               ret.DtDdosDstZoneIpProtoProtoNamePortIndOper.Oper.Sources,
		},
	}
}

func setSliceDdosDstZoneIpProtoProtoNamePortIndOperOperSrcEntryList(d []edpt.DdosDstZoneIpProtoProtoNamePortIndOperOperSrcEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["src_address_str"] = item.SrcAddressStr
		in["indicators"] = setSliceDdosDstZoneIpProtoProtoNamePortIndOperOperSrcEntryListIndicators(item.Indicators)
		in["detection_data_source"] = item.DetectionDataSource
		in["total_score"] = item.TotalScore
		in["current_level"] = item.CurrentLevel
		in["src_level"] = item.SrcLevel
		in["escalation_timestamp"] = item.EscalationTimestamp
		in["initial_learning"] = item.InitialLearning
		in["active_time"] = item.ActiveTime
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNamePortIndOperOperSrcEntryListIndicators(d []edpt.DdosDstZoneIpProtoProtoNamePortIndOperOperSrcEntryListIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["rate"] = item.Rate
		in["src_maximum"] = item.SrcMaximum
		in["src_minimum"] = item.SrcMinimum
		in["src_non_zero_minimum"] = item.SrcNonZeroMinimum
		in["src_average"] = item.SrcAverage
		in["score"] = item.Score
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNamePortIndOperOperIndicators(d []edpt.DdosDstZoneIpProtoProtoNamePortIndOperOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["rate"] = item.Rate
		in["zone_maximum"] = item.ZoneMaximum
		in["zone_minimum"] = item.ZoneMinimum
		in["zone_non_zero_minimum"] = item.ZoneNonZeroMinimum
		in["zone_average"] = item.ZoneAverage
		in["zone_adaptive_threshold"] = item.ZoneAdaptiveThreshold
		in["src_maximum"] = item.SrcMaximum
		in["indicator_cfg"] = setSliceDdosDstZoneIpProtoProtoNamePortIndOperOperIndicatorsIndicatorCfg(item.IndicatorCfg)
		in["score"] = item.Score
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNamePortIndOperOperIndicatorsIndicatorCfg(d []edpt.DdosDstZoneIpProtoProtoNamePortIndOperOperIndicatorsIndicatorCfg) []map[string]interface{} {
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

func getObjectDdosDstZoneIpProtoProtoNamePortIndOperOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNamePortIndOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNamePortIndOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcEntryList = getSliceDdosDstZoneIpProtoProtoNamePortIndOperOperSrcEntryList(in["src_entry_list"].([]interface{}))
		ret.Indicators = getSliceDdosDstZoneIpProtoProtoNamePortIndOperOperIndicators(in["indicators"].([]interface{}))
		ret.DetectionDataSource = in["detection_data_source"].(string)
		ret.TotalScore = in["total_score"].(string)
		ret.CurrentLevel = in["current_level"].(string)
		ret.EscalationTimestamp = in["escalation_timestamp"].(string)
		ret.InitialLearning = in["initial_learning"].(string)
		ret.ActiveTime = in["active_time"].(int)
		ret.SourcesAllEntries = in["sources_all_entries"].(int)
		ret.SubnetIpAddr = in["subnet_ip_addr"].(string)
		ret.SubnetIpv6Addr = in["subnet_ipv6_addr"].(string)
		ret.Ipv6 = in["ipv6"].(string)
		ret.Details = in["details"].(int)
		ret.Sources = in["sources"].(int)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNamePortIndOperOperSrcEntryList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNamePortIndOperOperSrcEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNamePortIndOperOperSrcEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNamePortIndOperOperSrcEntryList
		oi.SrcAddressStr = in["src_address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneIpProtoProtoNamePortIndOperOperSrcEntryListIndicators(in["indicators"].([]interface{}))
		oi.DetectionDataSource = in["detection_data_source"].(string)
		oi.TotalScore = in["total_score"].(string)
		oi.CurrentLevel = in["current_level"].(string)
		oi.SrcLevel = in["src_level"].(string)
		oi.EscalationTimestamp = in["escalation_timestamp"].(string)
		oi.InitialLearning = in["initial_learning"].(string)
		oi.ActiveTime = in["active_time"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNamePortIndOperOperSrcEntryListIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNamePortIndOperOperSrcEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNamePortIndOperOperSrcEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNamePortIndOperOperSrcEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.SrcMaximum = in["src_maximum"].(string)
		oi.SrcMinimum = in["src_minimum"].(string)
		oi.SrcNonZeroMinimum = in["src_non_zero_minimum"].(string)
		oi.SrcAverage = in["src_average"].(string)
		oi.Score = in["score"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNamePortIndOperOperIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNamePortIndOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNamePortIndOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNamePortIndOperOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.ZoneMaximum = in["zone_maximum"].(string)
		oi.ZoneMinimum = in["zone_minimum"].(string)
		oi.ZoneNonZeroMinimum = in["zone_non_zero_minimum"].(string)
		oi.ZoneAverage = in["zone_average"].(string)
		oi.ZoneAdaptiveThreshold = in["zone_adaptive_threshold"].(string)
		oi.SrcMaximum = in["src_maximum"].(string)
		oi.IndicatorCfg = getSliceDdosDstZoneIpProtoProtoNamePortIndOperOperIndicatorsIndicatorCfg(in["indicator_cfg"].([]interface{}))
		oi.Score = in["score"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNamePortIndOperOperIndicatorsIndicatorCfg(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNamePortIndOperOperIndicatorsIndicatorCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNamePortIndOperOperIndicatorsIndicatorCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNamePortIndOperOperIndicatorsIndicatorCfg
		oi.Level = in["level"].(int)
		oi.ZoneThreshold = in["zone_threshold"].(string)
		oi.SourceThreshold = in["source_threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoNamePortIndOper(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNamePortIndOper {
	var ret edpt.DdosDstZoneIpProtoProtoNamePortIndOper

	ret.Oper = getObjectDdosDstZoneIpProtoProtoNamePortIndOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
