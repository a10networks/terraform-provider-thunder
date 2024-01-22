package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneSrcPortRangePortIndOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_src_port_range_port_ind_oper`: Operational Status for the object port-ind\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneSrcPortRangePortIndOperRead,

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
									"rate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"indicator_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"level": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"zone_threshold": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"source_threshold": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"detection_data_source": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"current_level": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"details": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"src_port_range_end": {
				Type: schema.TypeString, Required: true, Description: "SrcPortRangeEnd",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"src_port_range_start": {
				Type: schema.TypeString, Required: true, Description: "SrcPortRangeStart",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
		},
	}
}

func resourceDdosDstZoneSrcPortRangePortIndOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortRangePortIndOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortRangePortIndOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneSrcPortRangePortIndOperOper := setObjectDdosDstZoneSrcPortRangePortIndOperOper(res)
		d.Set("oper", DdosDstZoneSrcPortRangePortIndOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneSrcPortRangePortIndOperOper(ret edpt.DataDdosDstZoneSrcPortRangePortIndOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators":            setSliceDdosDstZoneSrcPortRangePortIndOperOperIndicators(ret.DtDdosDstZoneSrcPortRangePortIndOper.Oper.Indicators),
			"detection_data_source": ret.DtDdosDstZoneSrcPortRangePortIndOper.Oper.DetectionDataSource,
			"current_level":         ret.DtDdosDstZoneSrcPortRangePortIndOper.Oper.CurrentLevel,
			"details":               ret.DtDdosDstZoneSrcPortRangePortIndOper.Oper.Details,
		},
	}
}

func setSliceDdosDstZoneSrcPortRangePortIndOperOperIndicators(d []edpt.DdosDstZoneSrcPortRangePortIndOperOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["rate"] = item.Rate
		in["indicator_cfg"] = setSliceDdosDstZoneSrcPortRangePortIndOperOperIndicatorsIndicatorCfg(item.IndicatorCfg)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneSrcPortRangePortIndOperOperIndicatorsIndicatorCfg(d []edpt.DdosDstZoneSrcPortRangePortIndOperOperIndicatorsIndicatorCfg) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["level"] = item.Level
		in["zone_threshold"] = item.ZoneThreshold
		in["source_threshold"] = item.SourceThreshold
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZoneSrcPortRangePortIndOperOper(d []interface{}) edpt.DdosDstZoneSrcPortRangePortIndOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortRangePortIndOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneSrcPortRangePortIndOperOperIndicators(in["indicators"].([]interface{}))
		ret.DetectionDataSource = in["detection_data_source"].(string)
		ret.CurrentLevel = in["current_level"].(string)
		ret.Details = in["details"].(int)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortRangePortIndOperOperIndicators(d []interface{}) []edpt.DdosDstZoneSrcPortRangePortIndOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortRangePortIndOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortRangePortIndOperOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.IndicatorCfg = getSliceDdosDstZoneSrcPortRangePortIndOperOperIndicatorsIndicatorCfg(in["indicator_cfg"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortRangePortIndOperOperIndicatorsIndicatorCfg(d []interface{}) []edpt.DdosDstZoneSrcPortRangePortIndOperOperIndicatorsIndicatorCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortRangePortIndOperOperIndicatorsIndicatorCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortRangePortIndOperOperIndicatorsIndicatorCfg
		oi.Level = in["level"].(int)
		oi.ZoneThreshold = in["zone_threshold"].(string)
		oi.SourceThreshold = in["source_threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneSrcPortRangePortIndOper(d *schema.ResourceData) edpt.DdosDstZoneSrcPortRangePortIndOper {
	var ret edpt.DdosDstZoneSrcPortRangePortIndOper

	ret.Oper = getObjectDdosDstZoneSrcPortRangePortIndOperOper(d.Get("oper").([]interface{}))

	ret.SrcPortRangeEnd = d.Get("src_port_range_end").(string)

	ret.ZoneName = d.Get("zone_name").(string)

	ret.SrcPortRangeStart = d.Get("src_port_range_start").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
