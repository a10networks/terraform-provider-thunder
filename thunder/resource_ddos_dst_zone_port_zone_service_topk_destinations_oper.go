package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceTopkDestinationsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_zone_service_topk_destinations_oper`: Operational Status for the object topk-destinations\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortZoneServiceTopkDestinationsOperRead,

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
			"port_num": {
				Type: schema.TypeString, Required: true, Description: "PortNum",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
		},
	}
}

func resourceDdosDstZonePortZoneServiceTopkDestinationsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceTopkDestinationsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceTopkDestinationsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortZoneServiceTopkDestinationsOperOper := setObjectDdosDstZonePortZoneServiceTopkDestinationsOperOper(res)
		d.Set("oper", DdosDstZonePortZoneServiceTopkDestinationsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortZoneServiceTopkDestinationsOperOper(ret edpt.DataDdosDstZonePortZoneServiceTopkDestinationsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators":     setSliceDdosDstZonePortZoneServiceTopkDestinationsOperOperIndicators(ret.DtDdosDstZonePortZoneServiceTopkDestinationsOper.Oper.Indicators),
			"entry_list":     setSliceDdosDstZonePortZoneServiceTopkDestinationsOperOperEntryList(ret.DtDdosDstZonePortZoneServiceTopkDestinationsOper.Oper.EntryList),
			"next_indicator": ret.DtDdosDstZonePortZoneServiceTopkDestinationsOper.Oper.NextIndicator,
			"finished":       ret.DtDdosDstZonePortZoneServiceTopkDestinationsOper.Oper.Finished,
			"details":        ret.DtDdosDstZonePortZoneServiceTopkDestinationsOper.Oper.Details,
			"top_k_key":      ret.DtDdosDstZonePortZoneServiceTopkDestinationsOper.Oper.TopKKey,
		},
	}
}

func setSliceDdosDstZonePortZoneServiceTopkDestinationsOperOperIndicators(d []edpt.DdosDstZonePortZoneServiceTopkDestinationsOperOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["destinations"] = setSliceDdosDstZonePortZoneServiceTopkDestinationsOperOperIndicatorsDestinations(item.Destinations)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceTopkDestinationsOperOperIndicatorsDestinations(d []edpt.DdosDstZonePortZoneServiceTopkDestinationsOperOperIndicatorsDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceTopkDestinationsOperOperEntryList(d []edpt.DdosDstZonePortZoneServiceTopkDestinationsOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZonePortZoneServiceTopkDestinationsOperOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceTopkDestinationsOperOperEntryListIndicators(d []edpt.DdosDstZonePortZoneServiceTopkDestinationsOperOperEntryListIndicators) []map[string]interface{} {
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

func getObjectDdosDstZonePortZoneServiceTopkDestinationsOperOper(d []interface{}) edpt.DdosDstZonePortZoneServiceTopkDestinationsOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceTopkDestinationsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZonePortZoneServiceTopkDestinationsOperOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZonePortZoneServiceTopkDestinationsOperOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceTopkDestinationsOperOperIndicators(d []interface{}) []edpt.DdosDstZonePortZoneServiceTopkDestinationsOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceTopkDestinationsOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceTopkDestinationsOperOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Destinations = getSliceDdosDstZonePortZoneServiceTopkDestinationsOperOperIndicatorsDestinations(in["destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceTopkDestinationsOperOperIndicatorsDestinations(d []interface{}) []edpt.DdosDstZonePortZoneServiceTopkDestinationsOperOperIndicatorsDestinations {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceTopkDestinationsOperOperIndicatorsDestinations, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceTopkDestinationsOperOperIndicatorsDestinations
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceTopkDestinationsOperOperEntryList(d []interface{}) []edpt.DdosDstZonePortZoneServiceTopkDestinationsOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceTopkDestinationsOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceTopkDestinationsOperOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZonePortZoneServiceTopkDestinationsOperOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceTopkDestinationsOperOperEntryListIndicators(d []interface{}) []edpt.DdosDstZonePortZoneServiceTopkDestinationsOperOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceTopkDestinationsOperOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceTopkDestinationsOperOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceTopkDestinationsOper(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceTopkDestinationsOper {
	var ret edpt.DdosDstZonePortZoneServiceTopkDestinationsOper

	ret.Oper = getObjectDdosDstZonePortZoneServiceTopkDestinationsOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.PortNum = d.Get("port_num").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
