package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneSrcPortZoneSrcPortPortIndOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_src_port_zone_src_port_port_ind_oper`: Operational Status for the object port-ind\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneSrcPortZoneSrcPortPortIndOperRead,

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

func resourceDdosDstZoneSrcPortZoneSrcPortPortIndOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortZoneSrcPortPortIndOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortZoneSrcPortPortIndOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneSrcPortZoneSrcPortPortIndOperOper := setObjectDdosDstZoneSrcPortZoneSrcPortPortIndOperOper(res)
		d.Set("oper", DdosDstZoneSrcPortZoneSrcPortPortIndOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneSrcPortZoneSrcPortPortIndOperOper(ret edpt.DataDdosDstZoneSrcPortZoneSrcPortPortIndOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators":            setSliceDdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicators(ret.DtDdosDstZoneSrcPortZoneSrcPortPortIndOper.Oper.Indicators),
			"detection_data_source": ret.DtDdosDstZoneSrcPortZoneSrcPortPortIndOper.Oper.DetectionDataSource,
			"current_level":         ret.DtDdosDstZoneSrcPortZoneSrcPortPortIndOper.Oper.CurrentLevel,
			"details":               ret.DtDdosDstZoneSrcPortZoneSrcPortPortIndOper.Oper.Details,
		},
	}
}

func setSliceDdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicators(d []edpt.DdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["rate"] = item.Rate
		in["indicator_cfg"] = setSliceDdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicatorsIndicatorCfg(item.IndicatorCfg)
		result = append(result, in)
	}
	return result
}

func setSliceDdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicatorsIndicatorCfg(d []edpt.DdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicatorsIndicatorCfg) []map[string]interface{} {
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

func getObjectDdosDstZoneSrcPortZoneSrcPortPortIndOperOper(d []interface{}) edpt.DdosDstZoneSrcPortZoneSrcPortPortIndOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneSrcPortZoneSrcPortPortIndOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicators(in["indicators"].([]interface{}))
		ret.DetectionDataSource = in["detection_data_source"].(string)
		ret.CurrentLevel = in["current_level"].(string)
		ret.Details = in["details"].(int)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicators(d []interface{}) []edpt.DdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.IndicatorCfg = getSliceDdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicatorsIndicatorCfg(in["indicator_cfg"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicatorsIndicatorCfg(d []interface{}) []edpt.DdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicatorsIndicatorCfg {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicatorsIndicatorCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicatorsIndicatorCfg
		oi.Level = in["level"].(int)
		oi.ZoneThreshold = in["zone_threshold"].(string)
		oi.SourceThreshold = in["source_threshold"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneSrcPortZoneSrcPortPortIndOper(d *schema.ResourceData) edpt.DdosDstZoneSrcPortZoneSrcPortPortIndOper {
	var ret edpt.DdosDstZoneSrcPortZoneSrcPortPortIndOper

	ret.Oper = getObjectDdosDstZoneSrcPortZoneSrcPortPortIndOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.PortNum = d.Get("port_num").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
