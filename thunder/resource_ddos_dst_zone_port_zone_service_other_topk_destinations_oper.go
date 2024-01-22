package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceOtherTopkDestinationsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_zone_service_other_topk_destinations_oper`: Operational Status for the object topk-destinations\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortZoneServiceOtherTopkDestinationsOperRead,

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

func resourceDdosDstZonePortZoneServiceOtherTopkDestinationsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceOtherTopkDestinationsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceOtherTopkDestinationsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortZoneServiceOtherTopkDestinationsOperOper := setObjectDdosDstZonePortZoneServiceOtherTopkDestinationsOperOper(res)
		d.Set("oper", DdosDstZonePortZoneServiceOtherTopkDestinationsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortZoneServiceOtherTopkDestinationsOperOper(ret edpt.DataDdosDstZonePortZoneServiceOtherTopkDestinationsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators":     setSliceDdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicators(ret.DtDdosDstZonePortZoneServiceOtherTopkDestinationsOper.Oper.Indicators),
			"entry_list":     setSliceDdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryList(ret.DtDdosDstZonePortZoneServiceOtherTopkDestinationsOper.Oper.EntryList),
			"next_indicator": ret.DtDdosDstZonePortZoneServiceOtherTopkDestinationsOper.Oper.NextIndicator,
			"finished":       ret.DtDdosDstZonePortZoneServiceOtherTopkDestinationsOper.Oper.Finished,
			"details":        ret.DtDdosDstZonePortZoneServiceOtherTopkDestinationsOper.Oper.Details,
			"top_k_key":      ret.DtDdosDstZonePortZoneServiceOtherTopkDestinationsOper.Oper.TopKKey,
		},
	}
}

func setSliceDdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicators(d []edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["destinations"] = setSliceDdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicatorsDestinations(item.Destinations)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicatorsDestinations(d []edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicatorsDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryList(d []edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryListIndicators(d []edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryListIndicators) []map[string]interface{} {
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

func getObjectDdosDstZonePortZoneServiceOtherTopkDestinationsOperOper(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicators(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Destinations = getSliceDdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicatorsDestinations(in["destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicatorsDestinations(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicatorsDestinations {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicatorsDestinations, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperIndicatorsDestinations
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryListIndicators(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOperOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceOtherTopkDestinationsOper(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOper {
	var ret edpt.DdosDstZonePortZoneServiceOtherTopkDestinationsOper

	ret.Oper = getObjectDdosDstZonePortZoneServiceOtherTopkDestinationsOperOper(d.Get("oper").([]interface{}))

	ret.PortOther = d.Get("port_other").(string)

	ret.ZoneName = d.Get("zone_name").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
