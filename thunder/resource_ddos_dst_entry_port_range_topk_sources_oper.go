package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryPortRangeTopkSourcesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_entry_port_range_topk_sources_oper`: Operational Status for the object topk-sources\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstEntryPortRangeTopkSourcesOperRead,

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
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"port_range_end": {
				Type: schema.TypeString, Required: true, Description: "PortRangeEnd",
			},
			"port_range_start": {
				Type: schema.TypeString, Required: true, Description: "PortRangeStart",
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}

func resourceDdosDstEntryPortRangeTopkSourcesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortRangeTopkSourcesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortRangeTopkSourcesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstEntryPortRangeTopkSourcesOperOper := setObjectDdosDstEntryPortRangeTopkSourcesOperOper(res)
		d.Set("oper", DdosDstEntryPortRangeTopkSourcesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstEntryPortRangeTopkSourcesOperOper(ret edpt.DataDdosDstEntryPortRangeTopkSourcesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators":     setSliceDdosDstEntryPortRangeTopkSourcesOperOperIndicators(ret.DtDdosDstEntryPortRangeTopkSourcesOper.Oper.Indicators),
			"entry_list":     setSliceDdosDstEntryPortRangeTopkSourcesOperOperEntryList(ret.DtDdosDstEntryPortRangeTopkSourcesOper.Oper.EntryList),
			"next_indicator": ret.DtDdosDstEntryPortRangeTopkSourcesOper.Oper.NextIndicator,
			"finished":       ret.DtDdosDstEntryPortRangeTopkSourcesOper.Oper.Finished,
			"details":        ret.DtDdosDstEntryPortRangeTopkSourcesOper.Oper.Details,
			"top_k_key":      ret.DtDdosDstEntryPortRangeTopkSourcesOper.Oper.TopKKey,
		},
	}
}

func setSliceDdosDstEntryPortRangeTopkSourcesOperOperIndicators(d []edpt.DdosDstEntryPortRangeTopkSourcesOperOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["sources"] = setSliceDdosDstEntryPortRangeTopkSourcesOperOperIndicatorsSources(item.Sources)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryPortRangeTopkSourcesOperOperIndicatorsSources(d []edpt.DdosDstEntryPortRangeTopkSourcesOperOperIndicatorsSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["rate"] = item.Rate
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryPortRangeTopkSourcesOperOperEntryList(d []edpt.DdosDstEntryPortRangeTopkSourcesOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address_str"] = item.AddressStr
		in["indicators"] = setSliceDdosDstEntryPortRangeTopkSourcesOperOperEntryListIndicators(item.Indicators)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstEntryPortRangeTopkSourcesOperOperEntryListIndicators(d []edpt.DdosDstEntryPortRangeTopkSourcesOperOperEntryListIndicators) []map[string]interface{} {
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

func getObjectDdosDstEntryPortRangeTopkSourcesOperOper(d []interface{}) edpt.DdosDstEntryPortRangeTopkSourcesOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeTopkSourcesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryPortRangeTopkSourcesOperOperIndicators(in["indicators"].([]interface{}))
		ret.EntryList = getSliceDdosDstEntryPortRangeTopkSourcesOperOperEntryList(in["entry_list"].([]interface{}))
		ret.NextIndicator = in["next_indicator"].(int)
		ret.Finished = in["finished"].(int)
		ret.Details = in["details"].(int)
		ret.TopKKey = in["top_k_key"].(string)
	}
	return ret
}

func getSliceDdosDstEntryPortRangeTopkSourcesOperOperIndicators(d []interface{}) []edpt.DdosDstEntryPortRangeTopkSourcesOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangeTopkSourcesOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangeTopkSourcesOperOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Sources = getSliceDdosDstEntryPortRangeTopkSourcesOperOperIndicatorsSources(in["sources"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryPortRangeTopkSourcesOperOperIndicatorsSources(d []interface{}) []edpt.DdosDstEntryPortRangeTopkSourcesOperOperIndicatorsSources {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangeTopkSourcesOperOperIndicatorsSources, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangeTopkSourcesOperOperIndicatorsSources
		oi.Address = in["address"].(string)
		oi.Rate = in["rate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryPortRangeTopkSourcesOperOperEntryList(d []interface{}) []edpt.DdosDstEntryPortRangeTopkSourcesOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangeTopkSourcesOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangeTopkSourcesOperOperEntryList
		oi.AddressStr = in["address_str"].(string)
		oi.Indicators = getSliceDdosDstEntryPortRangeTopkSourcesOperOperEntryListIndicators(in["indicators"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntryPortRangeTopkSourcesOperOperEntryListIndicators(d []interface{}) []edpt.DdosDstEntryPortRangeTopkSourcesOperOperEntryListIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangeTopkSourcesOperOperEntryListIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangeTopkSourcesOperOperEntryListIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.MaxPeak = in["max_peak"].(string)
		oi.PsdWdwCnt = in["psd_wdw_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstEntryPortRangeTopkSourcesOper(d *schema.ResourceData) edpt.DdosDstEntryPortRangeTopkSourcesOper {
	var ret edpt.DdosDstEntryPortRangeTopkSourcesOper

	ret.Oper = getObjectDdosDstEntryPortRangeTopkSourcesOperOper(d.Get("oper").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.PortRangeEnd = d.Get("port_range_end").(string)

	ret.PortRangeStart = d.Get("port_range_start").(string)

	ret.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
