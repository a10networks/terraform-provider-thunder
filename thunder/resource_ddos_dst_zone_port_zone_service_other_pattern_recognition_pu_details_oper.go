package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_zone_service_other_pattern_recognition_pu_details_oper`: Operational Status for the object pattern-recognition-pu-details\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all_filters": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"processing_unit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"timestamp": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"peace_pkt_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"war_pkt_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"war_pkt_percentage": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"filter_threshold": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"filter_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"filter_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"filter_enabled": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"hardware_filter": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"filter_expr": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"filter_desc": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"sample_ratio": {
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
			"port_other": {
				Type: schema.TypeString, Required: true, Description: "PortOther",
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

func resourceDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOper := setObjectDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOper(res)
		d.Set("oper", DdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOper(ret edpt.DataDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"all_filters": setSliceDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOperAllFilters(ret.DtDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOper.Oper.AllFilters),
		},
	}
}

func setSliceDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOperAllFilters(d []edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOperAllFilters) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["processing_unit"] = item.ProcessingUnit
		in["state"] = item.State
		in["timestamp"] = item.Timestamp
		in["peace_pkt_count"] = item.PeacePktCount
		in["war_pkt_count"] = item.WarPktCount
		in["war_pkt_percentage"] = item.WarPktPercentage
		in["filter_threshold"] = item.FilterThreshold
		in["filter_count"] = item.FilterCount
		in["filter_list"] = setSliceDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOperAllFiltersFilterList(item.FilterList)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOperAllFiltersFilterList(d []edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOperAllFiltersFilterList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["filter_enabled"] = item.FilterEnabled
		in["hardware_filter"] = item.HardwareFilter
		in["filter_expr"] = item.FilterExpr
		in["filter_desc"] = item.FilterDesc
		in["sample_ratio"] = item.SampleRatio
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOper(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllFilters = getSliceDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOperAllFilters(in["all_filters"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOperAllFilters(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOperAllFilters {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOperAllFilters, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOperAllFilters
		oi.ProcessingUnit = in["processing_unit"].(string)
		oi.State = in["state"].(string)
		oi.Timestamp = in["timestamp"].(string)
		oi.PeacePktCount = in["peace_pkt_count"].(int)
		oi.WarPktCount = in["war_pkt_count"].(int)
		oi.WarPktPercentage = in["war_pkt_percentage"].(int)
		oi.FilterThreshold = in["filter_threshold"].(int)
		oi.FilterCount = in["filter_count"].(int)
		oi.FilterList = getSliceDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOperAllFiltersFilterList(in["filter_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOperAllFiltersFilterList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOperAllFiltersFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOperAllFiltersFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOperAllFiltersFilterList
		oi.FilterEnabled = in["filter_enabled"].(int)
		oi.HardwareFilter = in["hardware_filter"].(int)
		oi.FilterExpr = in["filter_expr"].(string)
		oi.FilterDesc = in["filter_desc"].(string)
		oi.SampleRatio = in["sample_ratio"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOper(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOper {
	var ret edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOper

	ret.Oper = getObjectDdosDstZonePortZoneServiceOtherPatternRecognitionPuDetailsOperOper(d.Get("oper").([]interface{}))

	ret.PortOther = d.Get("port_other").(string)

	ret.ZoneName = d.Get("zone_name").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
