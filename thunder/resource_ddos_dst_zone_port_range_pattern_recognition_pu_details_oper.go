package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortRangePatternRecognitionPuDetailsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_range_pattern_recognition_pu_details_oper`: Operational Status for the object pattern-recognition-pu-details\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortRangePatternRecognitionPuDetailsOperRead,

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
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"port_range_end": {
				Type: schema.TypeString, Required: true, Description: "PortRangeEnd",
			},
			"port_range_start": {
				Type: schema.TypeString, Required: true, Description: "PortRangeStart",
			},
		},
	}
}

func resourceDdosDstZonePortRangePatternRecognitionPuDetailsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangePatternRecognitionPuDetailsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangePatternRecognitionPuDetailsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortRangePatternRecognitionPuDetailsOperOper := setObjectDdosDstZonePortRangePatternRecognitionPuDetailsOperOper(res)
		d.Set("oper", DdosDstZonePortRangePatternRecognitionPuDetailsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortRangePatternRecognitionPuDetailsOperOper(ret edpt.DataDdosDstZonePortRangePatternRecognitionPuDetailsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"all_filters": setSliceDdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFilters(ret.DtDdosDstZonePortRangePatternRecognitionPuDetailsOper.Oper.AllFilters),
		},
	}
}

func setSliceDdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFilters(d []edpt.DdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFilters) []map[string]interface{} {
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
		in["filter_list"] = setSliceDdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFiltersFilterList(item.FilterList)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFiltersFilterList(d []edpt.DdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFiltersFilterList) []map[string]interface{} {
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

func getObjectDdosDstZonePortRangePatternRecognitionPuDetailsOperOper(d []interface{}) edpt.DdosDstZonePortRangePatternRecognitionPuDetailsOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangePatternRecognitionPuDetailsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllFilters = getSliceDdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFilters(in["all_filters"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFilters(d []interface{}) []edpt.DdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFilters {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFilters, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFilters
		oi.ProcessingUnit = in["processing_unit"].(string)
		oi.State = in["state"].(string)
		oi.Timestamp = in["timestamp"].(string)
		oi.PeacePktCount = in["peace_pkt_count"].(int)
		oi.WarPktCount = in["war_pkt_count"].(int)
		oi.WarPktPercentage = in["war_pkt_percentage"].(int)
		oi.FilterThreshold = in["filter_threshold"].(int)
		oi.FilterCount = in["filter_count"].(int)
		oi.FilterList = getSliceDdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFiltersFilterList(in["filter_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFiltersFilterList(d []interface{}) []edpt.DdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFiltersFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFiltersFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangePatternRecognitionPuDetailsOperOperAllFiltersFilterList
		oi.FilterEnabled = in["filter_enabled"].(int)
		oi.HardwareFilter = in["hardware_filter"].(int)
		oi.FilterExpr = in["filter_expr"].(string)
		oi.FilterDesc = in["filter_desc"].(string)
		oi.SampleRatio = in["sample_ratio"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortRangePatternRecognitionPuDetailsOper(d *schema.ResourceData) edpt.DdosDstZonePortRangePatternRecognitionPuDetailsOper {
	var ret edpt.DdosDstZonePortRangePatternRecognitionPuDetailsOper

	ret.Oper = getObjectDdosDstZonePortRangePatternRecognitionPuDetailsOperOper(d.Get("oper").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.ZoneName = d.Get("zone_name").(string)

	ret.PortRangeEnd = d.Get("port_range_end").(string)

	ret.PortRangeStart = d.Get("port_range_start").(string)
	return ret
}
