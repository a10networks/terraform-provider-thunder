package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNumberTopkSourcesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_ip_proto_proto_number_topk_sources_oper`: Operational Status for the object topk-sources\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneIpProtoProtoNumberTopkSourcesOperRead,

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
			"protocol_num": {
				Type: schema.TypeString, Required: true, Description: "ProtocolNum",
			},
		},
	}
}

func resourceDdosDstZoneIpProtoProtoNumberTopkSourcesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberTopkSourcesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberTopkSourcesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneIpProtoProtoNumberTopkSourcesOperOper := setObjectDdosDstZoneIpProtoProtoNumberTopkSourcesOperOper(res)
		d.Set("oper", DdosDstZoneIpProtoProtoNumberTopkSourcesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneIpProtoProtoNumberTopkSourcesOperOper(ret edpt.DataDdosDstZoneIpProtoProtoNumberTopkSourcesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators":     setSliceDdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicators(ret.DtDdosDstZoneIpProtoProtoNumberTopkSourcesOper.Oper.Indicators),
			"entry_list":     setSliceDdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryList(ret.DtDdosDstZoneIpProtoProtoNumberTopkSourcesOper.Oper.EntryList),
			"next_indicator": ret.DtDdosDstZoneIpProtoProtoNumberTopkSourcesOper.Oper.NextIndicator,
			"finished":       ret.DtDdosDstZoneIpProtoProtoNumberTopkSourcesOper.Oper.Finished,
			"details":        ret.DtDdosDstZoneIpProtoProtoNumberTopkSourcesOper.Oper.Details,
			"top_k_key":      ret.DtDdosDstZoneIpProtoProtoNumberTopkSourcesOper.Oper.TopKKey,
		},
	}
}

func setSliceDdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicators(d []edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicatorsSources(d []edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryList(d []edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryListIndicators(d []edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryListIndicators) []map[string]interface{} {
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

func getObjectDdosDstZoneIpProtoProtoNumberTopkSourcesOperOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicatorsSources(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOperOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoNumberTopkSourcesOper(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOper {
	var ret edpt.DdosDstZoneIpProtoProtoNumberTopkSourcesOper

	ret.Oper = getObjectDdosDstZoneIpProtoProtoNumberTopkSourcesOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.ProtocolNum = d.Get("protocol_num").(string)
	return ret
}
