package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServicePatternRecognitionPuDetailsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_zone_service_pattern_recognition_pu_details_oper`: Operational Status for the object pattern-recognition-pu-details\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortZoneServicePatternRecognitionPuDetailsOperRead,

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
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"port_num": {
				Type: schema.TypeString, Required: true, Description: "PortNum",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
		},
	}
}

func resourceDdosDstZonePortZoneServicePatternRecognitionPuDetailsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServicePatternRecognitionPuDetailsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServicePatternRecognitionPuDetailsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOper := setObjectDdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOper(res)
		d.Set("oper", DdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOper(ret edpt.DataDdosDstZonePortZoneServicePatternRecognitionPuDetailsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"all_filters": setSliceDdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFilters(ret.DtDdosDstZonePortZoneServicePatternRecognitionPuDetailsOper.Oper.AllFilters),
		},
	}
}

func setSliceDdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFilters(d []edpt.DdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFilters) []map[string]interface{} {
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
		in["filter_list"] = setSliceDdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFiltersFilterList(item.FilterList)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFiltersFilterList(d []edpt.DdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFiltersFilterList) []map[string]interface{} {
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

func getObjectDdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOper(d []interface{}) edpt.DdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllFilters = getSliceDdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFilters(in["all_filters"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFilters(d []interface{}) []edpt.DdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFilters {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFilters, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFilters
		oi.ProcessingUnit = in["processing_unit"].(string)
		oi.State = in["state"].(string)
		oi.Timestamp = in["timestamp"].(string)
		oi.PeacePktCount = in["peace_pkt_count"].(int)
		oi.WarPktCount = in["war_pkt_count"].(int)
		oi.WarPktPercentage = in["war_pkt_percentage"].(int)
		oi.FilterThreshold = in["filter_threshold"].(int)
		oi.FilterCount = in["filter_count"].(int)
		oi.FilterList = getSliceDdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFiltersFilterList(in["filter_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFiltersFilterList(d []interface{}) []edpt.DdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFiltersFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFiltersFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOperAllFiltersFilterList
		oi.FilterEnabled = in["filter_enabled"].(int)
		oi.HardwareFilter = in["hardware_filter"].(int)
		oi.FilterExpr = in["filter_expr"].(string)
		oi.FilterDesc = in["filter_desc"].(string)
		oi.SampleRatio = in["sample_ratio"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServicePatternRecognitionPuDetailsOper(d *schema.ResourceData) edpt.DdosDstZonePortZoneServicePatternRecognitionPuDetailsOper {
	var ret edpt.DdosDstZonePortZoneServicePatternRecognitionPuDetailsOper

	ret.Oper = getObjectDdosDstZonePortZoneServicePatternRecognitionPuDetailsOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.PortNum = d.Get("port_num").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
