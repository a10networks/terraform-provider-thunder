package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneDetectionVictimIpDetectionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_detection_victim_ip_detection_oper`: Operational Status for the object victim-ip-detection\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneDetectionVictimIpDetectionOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address_str": {
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
														},
													},
												},
												"is_anomaly": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
									"is_learning_done": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"is_histogram_learning_done": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"is_ip_anomaly": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"is_static_threshold": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"escalation_timestamp": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"de_escalation_timestamp": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"ip_entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_ip_entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"active_list": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"victim_list": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipv4_ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6_ip": {
							Type: schema.TypeString, Optional: true, Description: "",
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

func resourceDdosDstZoneDetectionVictimIpDetectionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionVictimIpDetectionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionVictimIpDetectionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneDetectionVictimIpDetectionOperOper := setObjectDdosDstZoneDetectionVictimIpDetectionOperOper(res)
		d.Set("oper", DdosDstZoneDetectionVictimIpDetectionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneDetectionVictimIpDetectionOperOper(ret edpt.DataDdosDstZoneDetectionVictimIpDetectionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ip_entry_list":        setSliceDdosDstZoneDetectionVictimIpDetectionOperOperIpEntryList(ret.DtDdosDstZoneDetectionVictimIpDetectionOper.Oper.IpEntryList),
			"ip_entry_count":       ret.DtDdosDstZoneDetectionVictimIpDetectionOper.Oper.IpEntryCount,
			"total_ip_entry_count": ret.DtDdosDstZoneDetectionVictimIpDetectionOper.Oper.TotalIpEntryCount,
			"active_list":          ret.DtDdosDstZoneDetectionVictimIpDetectionOper.Oper.ActiveList,
			"victim_list":          ret.DtDdosDstZoneDetectionVictimIpDetectionOper.Oper.VictimList,
			"ipv4_ip":              ret.DtDdosDstZoneDetectionVictimIpDetectionOper.Oper.Ipv4Ip,
			"ipv6_ip":              ret.DtDdosDstZoneDetectionVictimIpDetectionOper.Oper.Ipv6Ip,
		},
	}
}

func setSliceDdosDstZoneDetectionVictimIpDetectionOperOperIpEntryList(d []edpt.DdosDstZoneDetectionVictimIpDetectionOperOperIpEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip_address_str"] = item.IpAddressStr
		in["indicators"] = setSliceDdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicators(item.Indicators)
		in["is_learning_done"] = item.IsLearningDone
		in["is_histogram_learning_done"] = item.IsHistogramLearningDone
		in["is_ip_anomaly"] = item.IsIpAnomaly
		in["is_static_threshold"] = item.Is_static_threshold
		in["escalation_timestamp"] = item.EscalationTimestamp
		in["de_escalation_timestamp"] = item.DeEscalationTimestamp
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicators(d []edpt.DdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["value"] = setSliceDdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicatorsValue(item.Value)
		in["is_anomaly"] = item.IsAnomaly
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicatorsValue(d []edpt.DdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicatorsValue) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["current"] = item.Current
		in["threshold"] = item.Threshold
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZoneDetectionVictimIpDetectionOperOper(d []interface{}) edpt.DdosDstZoneDetectionVictimIpDetectionOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneDetectionVictimIpDetectionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpEntryList = getSliceDdosDstZoneDetectionVictimIpDetectionOperOperIpEntryList(in["ip_entry_list"].([]interface{}))
		ret.IpEntryCount = in["ip_entry_count"].(int)
		ret.TotalIpEntryCount = in["total_ip_entry_count"].(int)
		ret.ActiveList = in["active_list"].(int)
		ret.VictimList = in["victim_list"].(int)
		ret.Ipv4Ip = in["ipv4_ip"].(string)
		ret.Ipv6Ip = in["ipv6_ip"].(string)
	}
	return ret
}

func getSliceDdosDstZoneDetectionVictimIpDetectionOperOperIpEntryList(d []interface{}) []edpt.DdosDstZoneDetectionVictimIpDetectionOperOperIpEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionVictimIpDetectionOperOperIpEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionVictimIpDetectionOperOperIpEntryList
		oi.IpAddressStr = in["ip_address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicators(in["indicators"].([]interface{}))
		oi.IsLearningDone = in["is_learning_done"].(int)
		oi.IsHistogramLearningDone = in["is_histogram_learning_done"].(int)
		oi.IsIpAnomaly = in["is_ip_anomaly"].(int)
		oi.Is_static_threshold = in["is_static_threshold"].(int)
		oi.EscalationTimestamp = in["escalation_timestamp"].(string)
		oi.DeEscalationTimestamp = in["de_escalation_timestamp"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicators(d []interface{}) []edpt.DdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Value = getSliceDdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicatorsValue(in["value"].([]interface{}))
		oi.IsAnomaly = in["is_anomaly"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicatorsValue(d []interface{}) []edpt.DdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicatorsValue {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicatorsValue, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionVictimIpDetectionOperOperIpEntryListIndicatorsValue
		oi.Current = in["current"].(string)
		oi.Threshold = in["threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneDetectionVictimIpDetectionOper(d *schema.ResourceData) edpt.DdosDstZoneDetectionVictimIpDetectionOper {
	var ret edpt.DdosDstZoneDetectionVictimIpDetectionOper

	ret.Oper = getObjectDdosDstZoneDetectionVictimIpDetectionOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)
	return ret
}
