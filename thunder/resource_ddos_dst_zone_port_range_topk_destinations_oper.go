package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortRangeTopkDestinationsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_range_topk_destinations_oper`: Operational Status for the object topk-destinations\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortRangeTopkDestinationsOperRead,

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

func resourceDdosDstZonePortRangeTopkDestinationsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeTopkDestinationsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeTopkDestinationsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortRangeTopkDestinationsOperOper := setObjectDdosDstZonePortRangeTopkDestinationsOperOper(res)
		d.Set("oper", DdosDstZonePortRangeTopkDestinationsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortRangeTopkDestinationsOperOper(ret edpt.DataDdosDstZonePortRangeTopkDestinationsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators":     setSliceDdosDstZonePortRangeTopkDestinationsOperOperIndicators(ret.DtDdosDstZonePortRangeTopkDestinationsOper.Oper.Indicators),
			"entry_list":     setSliceDdosDstZonePortRangeTopkDestinationsOperOperEntryList(ret.DtDdosDstZonePortRangeTopkDestinationsOper.Oper.EntryList),
			"next_indicator": ret.DtDdosDstZonePortRangeTopkDestinationsOper.Oper.NextIndicator,
			"finished":       ret.DtDdosDstZonePortRangeTopkDestinationsOper.Oper.Finished,
			"details":        ret.DtDdosDstZonePortRangeTopkDestinationsOper.Oper.Details,
			"top_k_key":      ret.DtDdosDstZonePortRangeTopkDestinationsOper.Oper.TopKKey,
		},
	}
}

func setSliceDdosDstZonePortRangeTopkDestinationsOperOperIndicators(d []edpt.DdosDstZonePortRangeTopkDestinationsOperOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["destinations"] = setSliceDdosDstZonePortRangeTopkDestinationsOperOperIndicatorsDestinations(item.Destinations)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortRangeTopkDestinationsOperOperIndicatorsDestinations(d []edpt.DdosDstZonePortRangeTopkDestinationsOperOperIndicatorsDestinations) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortRangeTopkDestinationsOperOperEntryList(d []edpt.DdosDstZonePortRangeTopkDestinationsOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZonePortRangeTopkDestinationsOperOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortRangeTopkDestinationsOperOperEntryListIndicators(d []edpt.DdosDstZonePortRangeTopkDestinationsOperOperEntryListIndicators) []map[string]interface{} {
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

func getObjectDdosDstZonePortRangeTopkDestinationsOperOper(d []interface{}) edpt.DdosDstZonePortRangeTopkDestinationsOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeTopkDestinationsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZonePortRangeTopkDestinationsOperOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZonePortRangeTopkDestinationsOperOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortRangeTopkDestinationsOperOperIndicators(d []interface{}) []edpt.DdosDstZonePortRangeTopkDestinationsOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeTopkDestinationsOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeTopkDestinationsOperOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Destinations = getSliceDdosDstZonePortRangeTopkDestinationsOperOperIndicatorsDestinations(in["destinations"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangeTopkDestinationsOperOperIndicatorsDestinations(d []interface{}) []edpt.DdosDstZonePortRangeTopkDestinationsOperOperIndicatorsDestinations {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeTopkDestinationsOperOperIndicatorsDestinations, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeTopkDestinationsOperOperIndicatorsDestinations
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangeTopkDestinationsOperOperEntryList(d []interface{}) []edpt.DdosDstZonePortRangeTopkDestinationsOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeTopkDestinationsOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeTopkDestinationsOperOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZonePortRangeTopkDestinationsOperOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortRangeTopkDestinationsOperOperEntryListIndicators(d []interface{}) []edpt.DdosDstZonePortRangeTopkDestinationsOperOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeTopkDestinationsOperOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeTopkDestinationsOperOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortRangeTopkDestinationsOper(d *schema.ResourceData) edpt.DdosDstZonePortRangeTopkDestinationsOper {
	var ret edpt.DdosDstZonePortRangeTopkDestinationsOper

	ret.Oper = getObjectDdosDstZonePortRangeTopkDestinationsOperOper(d.Get("oper").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.ZoneName = d.Get("zone_name").(string)

	ret.PortRangeEnd = d.Get("port_range_end").(string)

	ret.PortRangeStart = d.Get("port_range_start").(string)
	return ret
}
