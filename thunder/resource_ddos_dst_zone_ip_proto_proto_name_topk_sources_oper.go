package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNameTopkSourcesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_ip_proto_proto_name_topk_sources_oper`: Operational Status for the object topk-sources\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneIpProtoProtoNameTopkSourcesOperRead,

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
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
		},
	}
}

func resourceDdosDstZoneIpProtoProtoNameTopkSourcesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameTopkSourcesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameTopkSourcesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneIpProtoProtoNameTopkSourcesOperOper := setObjectDdosDstZoneIpProtoProtoNameTopkSourcesOperOper(res)
		d.Set("oper", DdosDstZoneIpProtoProtoNameTopkSourcesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneIpProtoProtoNameTopkSourcesOperOper(ret edpt.DataDdosDstZoneIpProtoProtoNameTopkSourcesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators":     setSliceDdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicators(ret.DtDdosDstZoneIpProtoProtoNameTopkSourcesOper.Oper.Indicators),
			"entry_list":     setSliceDdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryList(ret.DtDdosDstZoneIpProtoProtoNameTopkSourcesOper.Oper.EntryList),
			"next_indicator": ret.DtDdosDstZoneIpProtoProtoNameTopkSourcesOper.Oper.NextIndicator,
			"finished":       ret.DtDdosDstZoneIpProtoProtoNameTopkSourcesOper.Oper.Finished,
			"details":        ret.DtDdosDstZoneIpProtoProtoNameTopkSourcesOper.Oper.Details,
			"top_k_key":      ret.DtDdosDstZoneIpProtoProtoNameTopkSourcesOper.Oper.TopKKey,
		},
	}
}

func setSliceDdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicators(d []edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicatorsSources(d []edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryList(d []edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryListIndicators(d []edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryListIndicators) []map[string]interface{} {
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

func getObjectDdosDstZoneIpProtoProtoNameTopkSourcesOperOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicatorsSources(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOperOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryListIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOperOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoNameTopkSourcesOper(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOper {
	var ret edpt.DdosDstZoneIpProtoProtoNameTopkSourcesOper

	ret.Oper = getObjectDdosDstZoneIpProtoProtoNameTopkSourcesOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
