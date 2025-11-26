package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortRangeProgressionTrackingOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_range_progression_tracking_oper`: Operational Status for the object progression-tracking\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortRangeProgressionTrackingOperRead,

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
									"num_sample": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"average": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"maximum": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"minimum": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"standard_deviation": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"learning_details": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"learning_brief": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"recommended_template": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"template_debug_table": {
							Type: schema.TypeInt, Optional: true, Description: "",
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

func resourceDdosDstZonePortRangeProgressionTrackingOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeProgressionTrackingOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeProgressionTrackingOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortRangeProgressionTrackingOperOper := setObjectDdosDstZonePortRangeProgressionTrackingOperOper(res)
		d.Set("oper", DdosDstZonePortRangeProgressionTrackingOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortRangeProgressionTrackingOperOper(ret edpt.DataDdosDstZonePortRangeProgressionTrackingOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators":           setSliceDdosDstZonePortRangeProgressionTrackingOperOperIndicators(ret.DtDdosDstZonePortRangeProgressionTrackingOper.Oper.Indicators),
			"learning_details":     ret.DtDdosDstZonePortRangeProgressionTrackingOper.Oper.LearningDetails,
			"learning_brief":       ret.DtDdosDstZonePortRangeProgressionTrackingOper.Oper.LearningBrief,
			"recommended_template": ret.DtDdosDstZonePortRangeProgressionTrackingOper.Oper.RecommendedTemplate,
			"template_debug_table": ret.DtDdosDstZonePortRangeProgressionTrackingOper.Oper.TemplateDebugTable,
		},
	}
}

func setSliceDdosDstZonePortRangeProgressionTrackingOperOperIndicators(d []edpt.DdosDstZonePortRangeProgressionTrackingOperOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["num_sample"] = item.NumSample
		in["average"] = item.Average
		in["maximum"] = item.Maximum
		in["minimum"] = item.Minimum
		in["standard_deviation"] = item.StandardDeviation
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZonePortRangeProgressionTrackingOperOper(d []interface{}) edpt.DdosDstZonePortRangeProgressionTrackingOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeProgressionTrackingOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZonePortRangeProgressionTrackingOperOperIndicators(in["indicators"].([]interface{}))
		ret.LearningDetails = in["learning_details"].(int)
		ret.LearningBrief = in["learning_brief"].(int)
		ret.RecommendedTemplate = in["recommended_template"].(int)
		ret.TemplateDebugTable = in["template_debug_table"].(int)
	}
	return ret
}

func getSliceDdosDstZonePortRangeProgressionTrackingOperOperIndicators(d []interface{}) []edpt.DdosDstZonePortRangeProgressionTrackingOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeProgressionTrackingOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeProgressionTrackingOperOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.NumSample = in["num_sample"].(int)
		oi.Average = in["average"].(string)
		oi.Maximum = in["maximum"].(string)
		oi.Minimum = in["minimum"].(string)
		oi.StandardDeviation = in["standard_deviation"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortRangeProgressionTrackingOper(d *schema.ResourceData) edpt.DdosDstZonePortRangeProgressionTrackingOper {
	var ret edpt.DdosDstZonePortRangeProgressionTrackingOper

	ret.Oper = getObjectDdosDstZonePortRangeProgressionTrackingOperOper(d.Get("oper").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.ZoneName = d.Get("zone_name").(string)

	ret.PortRangeEnd = d.Get("port_range_end").(string)

	ret.PortRangeStart = d.Get("port_range_start").(string)
	return ret
}
