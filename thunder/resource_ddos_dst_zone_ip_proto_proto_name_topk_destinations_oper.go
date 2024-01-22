package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNameTopkDestinationsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_ip_proto_proto_name_topk_destinations_oper`: Operational Status for the object topk-destinations\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneIpProtoProtoNameTopkDestinationsOperRead,

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
						"entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address_str": {
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
												"max_peak": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"psd_wdw_cnt": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"next_indicator": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"finished": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"details": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"top_k_key": {
							Type: schema.TypeString, Optional: true, Description: "",
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

func resourceDdosDstZoneIpProtoProtoNameTopkDestinationsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameTopkDestinationsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameTopkDestinationsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneIpProtoProtoNameTopkDestinationsOperOper := setObjectDdosDstZoneIpProtoProtoNameTopkDestinationsOperOper(res)
		d.Set("oper", DdosDstZoneIpProtoProtoNameTopkDestinationsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneIpProtoProtoNameTopkDestinationsOperOper(ret edpt.DataDdosDstZoneIpProtoProtoNameTopkDestinationsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators":     setSliceDdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicators(ret.DtDdosDstZoneIpProtoProtoNameTopkDestinationsOper.Oper.Indicators),
			"entry_list":     setSliceDdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryList(ret.DtDdosDstZoneIpProtoProtoNameTopkDestinationsOper.Oper.EntryList),
			"next_indicator": ret.DtDdosDstZoneIpProtoProtoNameTopkDestinationsOper.Oper.NextIndicator,
			"finished":       ret.DtDdosDstZoneIpProtoProtoNameTopkDestinationsOper.Oper.Finished,
			"details":        ret.DtDdosDstZoneIpProtoProtoNameTopkDestinationsOper.Oper.Details,
			"top_k_key":      ret.DtDdosDstZoneIpProtoProtoNameTopkDestinationsOper.Oper.TopKKey,
		},
	}
}

func setSliceDdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicators(d []edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["destinations"] = setSliceDdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicatorsDestinations(item.Destinations)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicatorsDestinations(d []edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicatorsDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryList(d []edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryListIndicators(d []edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryListIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["rate"] = item.Rate
		in["max_peak"] = item.MaxPeak
		in["psd_wdw_cnt"] = item.PsdWdwCnt
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZoneIpProtoProtoNameTopkDestinationsOperOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Destinations = getSliceDdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicatorsDestinations(in["destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicatorsDestinations(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicatorsDestinations {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicatorsDestinations, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperIndicatorsDestinations
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOperOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoNameTopkDestinationsOper(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOper {
	var ret edpt.DdosDstZoneIpProtoProtoNameTopkDestinationsOper

	ret.Oper = getObjectDdosDstZoneIpProtoProtoNameTopkDestinationsOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
