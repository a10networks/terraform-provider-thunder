package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceTopkSourcesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_zone_service_topk_sources_oper`: Operational Status for the object topk-sources\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortZoneServiceTopkSourcesOperRead,

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

func resourceDdosDstZonePortZoneServiceTopkSourcesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceTopkSourcesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceTopkSourcesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortZoneServiceTopkSourcesOperOper := setObjectDdosDstZonePortZoneServiceTopkSourcesOperOper(res)
		d.Set("oper", DdosDstZonePortZoneServiceTopkSourcesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortZoneServiceTopkSourcesOperOper(ret edpt.DataDdosDstZonePortZoneServiceTopkSourcesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators":     setSliceDdosDstZonePortZoneServiceTopkSourcesOperOperIndicators(ret.DtDdosDstZonePortZoneServiceTopkSourcesOper.Oper.Indicators),
			"entry_list":     setSliceDdosDstZonePortZoneServiceTopkSourcesOperOperEntryList(ret.DtDdosDstZonePortZoneServiceTopkSourcesOper.Oper.EntryList),
			"next_indicator": ret.DtDdosDstZonePortZoneServiceTopkSourcesOper.Oper.NextIndicator,
			"finished":       ret.DtDdosDstZonePortZoneServiceTopkSourcesOper.Oper.Finished,
			"details":        ret.DtDdosDstZonePortZoneServiceTopkSourcesOper.Oper.Details,
			"top_k_key":      ret.DtDdosDstZonePortZoneServiceTopkSourcesOper.Oper.TopKKey,
		},
	}
}

func setSliceDdosDstZonePortZoneServiceTopkSourcesOperOperIndicators(d []edpt.DdosDstZonePortZoneServiceTopkSourcesOperOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstZonePortZoneServiceTopkSourcesOperOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceTopkSourcesOperOperIndicatorsSources(d []edpt.DdosDstZonePortZoneServiceTopkSourcesOperOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceTopkSourcesOperOperEntryList(d []edpt.DdosDstZonePortZoneServiceTopkSourcesOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZonePortZoneServiceTopkSourcesOperOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZonePortZoneServiceTopkSourcesOperOperEntryListIndicators(d []edpt.DdosDstZonePortZoneServiceTopkSourcesOperOperEntryListIndicators) []map[string]interface{} {
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

func getObjectDdosDstZonePortZoneServiceTopkSourcesOperOper(d []interface{}) edpt.DdosDstZonePortZoneServiceTopkSourcesOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceTopkSourcesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZonePortZoneServiceTopkSourcesOperOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZonePortZoneServiceTopkSourcesOperOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceTopkSourcesOperOperIndicators(d []interface{}) []edpt.DdosDstZonePortZoneServiceTopkSourcesOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceTopkSourcesOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceTopkSourcesOperOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstZonePortZoneServiceTopkSourcesOperOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceTopkSourcesOperOperIndicatorsSources(d []interface{}) []edpt.DdosDstZonePortZoneServiceTopkSourcesOperOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceTopkSourcesOperOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceTopkSourcesOperOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceTopkSourcesOperOperEntryList(d []interface{}) []edpt.DdosDstZonePortZoneServiceTopkSourcesOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceTopkSourcesOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceTopkSourcesOperOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZonePortZoneServiceTopkSourcesOperOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceTopkSourcesOperOperEntryListIndicators(d []interface{}) []edpt.DdosDstZonePortZoneServiceTopkSourcesOperOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceTopkSourcesOperOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceTopkSourcesOperOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceTopkSourcesOper(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceTopkSourcesOper {
	var ret edpt.DdosDstZonePortZoneServiceTopkSourcesOper

	ret.Oper = getObjectDdosDstZonePortZoneServiceTopkSourcesOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.PortNum = d.Get("port_num").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
