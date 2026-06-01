package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNameProgressionTrackingOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_ip_proto_proto_name_progression_tracking_oper`: Operational Status for the object progression-tracking\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneIpProtoProtoNameProgressionTrackingOperRead,

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
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
		},
	}
}

func resourceDdosDstZoneIpProtoProtoNameProgressionTrackingOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameProgressionTrackingOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameProgressionTrackingOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneIpProtoProtoNameProgressionTrackingOperOper := setObjectDdosDstZoneIpProtoProtoNameProgressionTrackingOperOper(res)
		d.Set("oper", DdosDstZoneIpProtoProtoNameProgressionTrackingOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneIpProtoProtoNameProgressionTrackingOperOper(ret edpt.DataDdosDstZoneIpProtoProtoNameProgressionTrackingOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators":           setSliceDdosDstZoneIpProtoProtoNameProgressionTrackingOperOperIndicators(ret.DtDdosDstZoneIpProtoProtoNameProgressionTrackingOper.Oper.Indicators),
			"learning_details":     ret.DtDdosDstZoneIpProtoProtoNameProgressionTrackingOper.Oper.LearningDetails,
			"learning_brief":       ret.DtDdosDstZoneIpProtoProtoNameProgressionTrackingOper.Oper.LearningBrief,
			"recommended_template": ret.DtDdosDstZoneIpProtoProtoNameProgressionTrackingOper.Oper.RecommendedTemplate,
			"template_debug_table": ret.DtDdosDstZoneIpProtoProtoNameProgressionTrackingOper.Oper.TemplateDebugTable,
		},
	}
}

func setSliceDdosDstZoneIpProtoProtoNameProgressionTrackingOperOperIndicators(d []edpt.DdosDstZoneIpProtoProtoNameProgressionTrackingOperOperIndicators) []map[string]interface{} {
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

func getObjectDdosDstZoneIpProtoProtoNameProgressionTrackingOperOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameProgressionTrackingOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameProgressionTrackingOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneIpProtoProtoNameProgressionTrackingOperOperIndicators(in["indicators"].([]interface{}))
		ret.LearningDetails = in["learning_details"].(int)
		ret.LearningBrief = in["learning_brief"].(int)
		ret.RecommendedTemplate = in["recommended_template"].(int)
		ret.TemplateDebugTable = in["template_debug_table"].(int)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameProgressionTrackingOperOperIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameProgressionTrackingOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameProgressionTrackingOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameProgressionTrackingOperOperIndicators
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

func dataToEndpointDdosDstZoneIpProtoProtoNameProgressionTrackingOper(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNameProgressionTrackingOper {
	var ret edpt.DdosDstZoneIpProtoProtoNameProgressionTrackingOper

	ret.Oper = getObjectDdosDstZoneIpProtoProtoNameProgressionTrackingOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
