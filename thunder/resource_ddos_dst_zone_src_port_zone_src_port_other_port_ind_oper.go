package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneSrcPortZoneSrcPortOtherPortIndOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_src_port_zone_src_port_other_port_ind_oper`: Operational Status for the object port-ind\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneSrcPortZoneSrcPortOtherPortIndOperRead,

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

func resourceDdosDstZoneSrcPortZoneSrcPortOtherPortIndOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortZoneSrcPortOtherPortIndOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortZoneSrcPortOtherPortIndOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOper := setObjectDdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOper(res)
		d.Set("oper", DdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOper(ret edpt.DataDdosDstZoneSrcPortZoneSrcPortOtherPortIndOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators":            setSliceDdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicators(ret.DtDdosDstZoneSrcPortZoneSrcPortOtherPortIndOper.Oper.Indicators),
			"detection_data_source": ret.DtDdosDstZoneSrcPortZoneSrcPortOtherPortIndOper.Oper.DetectionDataSource,
			"current_level":         ret.DtDdosDstZoneSrcPortZoneSrcPortOtherPortIndOper.Oper.CurrentLevel,
			"details":               ret.DtDdosDstZoneSrcPortZoneSrcPortOtherPortIndOper.Oper.Details,
		},
	}
}

func setSliceDdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicators(d []edpt.DdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["rate"] = item.Rate
		in["indicator_cfg"] = setSliceDdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicatorsIndicatorCfg(item.IndicatorCfg)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicatorsIndicatorCfg(d []edpt.DdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicatorsIndicatorCfg) []map[string]interface{} {
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

func getObjectDdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOper(d []interface{}) edpt.DdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicators(in["indicators"].([]interface{}))
		ret.DetectionDataSource = in["detection_data_source"].(string)
		ret.CurrentLevel = in["current_level"].(string)
		ret.Details = in["details"].(int)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicators(d []interface{}) []edpt.DdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.IndicatorCfg = getSliceDdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicatorsIndicatorCfg(in["indicator_cfg"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicatorsIndicatorCfg(d []interface{}) []edpt.DdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicatorsIndicatorCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicatorsIndicatorCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicatorsIndicatorCfg
		oi.Level = in["level"].(int)
		oi.ZoneThreshold = in["zone_threshold"].(string)
		oi.SourceThreshold = in["source_threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneSrcPortZoneSrcPortOtherPortIndOper(d *schema.ResourceData) edpt.DdosDstZoneSrcPortZoneSrcPortOtherPortIndOper {
	var ret edpt.DdosDstZoneSrcPortZoneSrcPortOtherPortIndOper

	ret.Oper = getObjectDdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOper(d.Get("oper").([]interface{}))

	ret.PortOther = d.Get("port_other").(string)

	ret.ZoneName = d.Get("zone_name").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
