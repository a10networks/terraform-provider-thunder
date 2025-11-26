package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_detection_outbound_detection_topk_source_subnet_oper`: Operational Status for the object topk-source-subnet\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperRead,

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
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}

func resourceDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOper := setObjectDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOper(res)
		d.Set("oper", DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOper(ret edpt.DataDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entry_list": setSliceDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryList(ret.DtDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOper.Oper.EntryList),
		},
	}
}

func setSliceDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryList(d []edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["location_type"] = item.LocationType
		in["location_name"] = item.LocationName
		in["indicators"] = setSliceDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicators(d []edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["source_subnets"] = setSliceDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicatorsSourceSubnets(item.SourceSubnets)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicatorsSourceSubnets(d []edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicatorsSourceSubnets) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOper(d []interface{}) edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EntryList = getSliceDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryList(in["entry_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryList(d []interface{}) []edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryList
		oi.LocationType = in["location_type"].(string)
		oi.LocationName = in["location_name"].(string)
		oi.Indicators = getSliceDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.SourceSubnets = getSliceDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicatorsSourceSubnets(in["source_subnets"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicatorsSourceSubnets(d []interface{}) []edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicatorsSourceSubnets {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicatorsSourceSubnets, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOperEntryListIndicatorsSourceSubnets
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOper(d *schema.ResourceData) edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOper {
	var ret edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOper

	ret.Oper = getObjectDdosDstZoneDetectionOutboundDetectionTopkSourceSubnetOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)
	return ret
}
