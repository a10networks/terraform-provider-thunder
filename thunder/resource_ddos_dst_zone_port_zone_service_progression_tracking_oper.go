package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceProgressionTrackingOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_zone_service_progression_tracking_oper`: Operational Status for the object progression-tracking\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortZoneServiceProgressionTrackingOperRead,

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

func resourceDdosDstZonePortZoneServiceProgressionTrackingOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceProgressionTrackingOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceProgressionTrackingOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortZoneServiceProgressionTrackingOperOper := setObjectDdosDstZonePortZoneServiceProgressionTrackingOperOper(res)
		d.Set("oper", DdosDstZonePortZoneServiceProgressionTrackingOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortZoneServiceProgressionTrackingOperOper(ret edpt.DataDdosDstZonePortZoneServiceProgressionTrackingOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators":           setSliceDdosDstZonePortZoneServiceProgressionTrackingOperOperIndicators(ret.DtDdosDstZonePortZoneServiceProgressionTrackingOper.Oper.Indicators),
			"learning_details":     ret.DtDdosDstZonePortZoneServiceProgressionTrackingOper.Oper.LearningDetails,
			"learning_brief":       ret.DtDdosDstZonePortZoneServiceProgressionTrackingOper.Oper.LearningBrief,
			"recommended_template": ret.DtDdosDstZonePortZoneServiceProgressionTrackingOper.Oper.RecommendedTemplate,
			"template_debug_table": ret.DtDdosDstZonePortZoneServiceProgressionTrackingOper.Oper.TemplateDebugTable,
		},
	}
}

func setSliceDdosDstZonePortZoneServiceProgressionTrackingOperOperIndicators(d []edpt.DdosDstZonePortZoneServiceProgressionTrackingOperOperIndicators) []map[string]interface{} {
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

func getObjectDdosDstZonePortZoneServiceProgressionTrackingOperOper(d []interface{}) edpt.DdosDstZonePortZoneServiceProgressionTrackingOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceProgressionTrackingOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZonePortZoneServiceProgressionTrackingOperOperIndicators(in["indicators"].([]interface{}))
		ret.LearningDetails = in["learning_details"].(int)
		ret.LearningBrief = in["learning_brief"].(int)
		ret.RecommendedTemplate = in["recommended_template"].(int)
		ret.TemplateDebugTable = in["template_debug_table"].(int)
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceProgressionTrackingOperOperIndicators(d []interface{}) []edpt.DdosDstZonePortZoneServiceProgressionTrackingOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceProgressionTrackingOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceProgressionTrackingOperOperIndicators
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

func dataToEndpointDdosDstZonePortZoneServiceProgressionTrackingOper(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceProgressionTrackingOper {
	var ret edpt.DdosDstZonePortZoneServiceProgressionTrackingOper

	ret.Oper = getObjectDdosDstZonePortZoneServiceProgressionTrackingOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.PortNum = d.Get("port_num").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
