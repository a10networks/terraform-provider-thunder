package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneTopkDestinationsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_topk_destinations_oper`: Operational Status for the object topk-destinations\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneTopkDestinationsOperRead,

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
		},
	}
}

func resourceDdosDstZoneTopkDestinationsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneTopkDestinationsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneTopkDestinationsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneTopkDestinationsOperOper := setObjectDdosDstZoneTopkDestinationsOperOper(res)
		d.Set("oper", DdosDstZoneTopkDestinationsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneTopkDestinationsOperOper(ret edpt.DataDdosDstZoneTopkDestinationsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators":     setSliceDdosDstZoneTopkDestinationsOperOperIndicators(ret.DtDdosDstZoneTopkDestinationsOper.Oper.Indicators),
			"entry_list":     setSliceDdosDstZoneTopkDestinationsOperOperEntryList(ret.DtDdosDstZoneTopkDestinationsOper.Oper.EntryList),
			"next_indicator": ret.DtDdosDstZoneTopkDestinationsOper.Oper.NextIndicator,
			"finished":       ret.DtDdosDstZoneTopkDestinationsOper.Oper.Finished,
			"details":        ret.DtDdosDstZoneTopkDestinationsOper.Oper.Details,
			"top_k_key":      ret.DtDdosDstZoneTopkDestinationsOper.Oper.TopKKey,
		},
	}
}

func setSliceDdosDstZoneTopkDestinationsOperOperIndicators(d []edpt.DdosDstZoneTopkDestinationsOperOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["destinations"] = setSliceDdosDstZoneTopkDestinationsOperOperIndicatorsDestinations(item.Destinations)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneTopkDestinationsOperOperIndicatorsDestinations(d []edpt.DdosDstZoneTopkDestinationsOperOperIndicatorsDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneTopkDestinationsOperOperEntryList(d []edpt.DdosDstZoneTopkDestinationsOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZoneTopkDestinationsOperOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneTopkDestinationsOperOperEntryListIndicators(d []edpt.DdosDstZoneTopkDestinationsOperOperEntryListIndicators) []map[string]interface{} {
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

func getObjectDdosDstZoneTopkDestinationsOperOper(d []interface{}) edpt.DdosDstZoneTopkDestinationsOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneTopkDestinationsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneTopkDestinationsOperOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZoneTopkDestinationsOperOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZoneTopkDestinationsOperOperIndicators(d []interface{}) []edpt.DdosDstZoneTopkDestinationsOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneTopkDestinationsOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneTopkDestinationsOperOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Destinations = getSliceDdosDstZoneTopkDestinationsOperOperIndicatorsDestinations(in["destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneTopkDestinationsOperOperIndicatorsDestinations(d []interface{}) []edpt.DdosDstZoneTopkDestinationsOperOperIndicatorsDestinations {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneTopkDestinationsOperOperIndicatorsDestinations, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneTopkDestinationsOperOperIndicatorsDestinations
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneTopkDestinationsOperOperEntryList(d []interface{}) []edpt.DdosDstZoneTopkDestinationsOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneTopkDestinationsOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneTopkDestinationsOperOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneTopkDestinationsOperOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneTopkDestinationsOperOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneTopkDestinationsOperOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneTopkDestinationsOperOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneTopkDestinationsOperOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneTopkDestinationsOper(d *schema.ResourceData) edpt.DdosDstZoneTopkDestinationsOper {
	var ret edpt.DdosDstZoneTopkDestinationsOper

	ret.Oper = getObjectDdosDstZoneTopkDestinationsOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)
	return ret
}
