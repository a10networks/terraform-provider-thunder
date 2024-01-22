package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceOtherTopkSourcesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_zone_service_other_topk_sources_oper`: Operational Status for the object topk-sources\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortZoneServiceOtherTopkSourcesOperRead,

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
									"sources": {
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

func resourceDdosDstZonePortZoneServiceOtherTopkSourcesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceOtherTopkSourcesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceOtherTopkSourcesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortZoneServiceOtherTopkSourcesOperOper := setObjectDdosDstZonePortZoneServiceOtherTopkSourcesOperOper(res)
		d.Set("oper", DdosDstZonePortZoneServiceOtherTopkSourcesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortZoneServiceOtherTopkSourcesOperOper(ret edpt.DataDdosDstZonePortZoneServiceOtherTopkSourcesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators":     setSliceDdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicators(ret.DtDdosDstZonePortZoneServiceOtherTopkSourcesOper.Oper.Indicators),
			"entry_list":     setSliceDdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryList(ret.DtDdosDstZonePortZoneServiceOtherTopkSourcesOper.Oper.EntryList),
			"next_indicator": ret.DtDdosDstZonePortZoneServiceOtherTopkSourcesOper.Oper.NextIndicator,
			"finished":       ret.DtDdosDstZonePortZoneServiceOtherTopkSourcesOper.Oper.Finished,
			"details":        ret.DtDdosDstZonePortZoneServiceOtherTopkSourcesOper.Oper.Details,
			"top_k_key":      ret.DtDdosDstZonePortZoneServiceOtherTopkSourcesOper.Oper.TopKKey,
		},
	}
}

func setSliceDdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicators(d []edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicatorsSources(d []edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryList(d []edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryListIndicators(d []edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryListIndicators) []map[string]interface{} {
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

func getObjectDdosDstZonePortZoneServiceOtherTopkSourcesOperOper(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicators(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicatorsSources(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOperOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryListIndicators(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOperOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceOtherTopkSourcesOper(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOper {
	var ret edpt.DdosDstZonePortZoneServiceOtherTopkSourcesOper

	ret.Oper = getObjectDdosDstZonePortZoneServiceOtherTopkSourcesOperOper(d.Get("oper").([]interface{}))

	ret.PortOther = d.Get("port_other").(string)

	ret.ZoneName = d.Get("zone_name").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
