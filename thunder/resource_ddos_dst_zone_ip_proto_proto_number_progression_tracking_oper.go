package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNumberProgressionTrackingOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_ip_proto_proto_number_progression_tracking_oper`: Operational Status for the object progression-tracking\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneIpProtoProtoNumberProgressionTrackingOperRead,

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
			"protocol_num": {
				Type: schema.TypeString, Required: true, Description: "ProtocolNum",
			},
		},
	}
}

func resourceDdosDstZoneIpProtoProtoNumberProgressionTrackingOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberProgressionTrackingOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberProgressionTrackingOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneIpProtoProtoNumberProgressionTrackingOperOper := setObjectDdosDstZoneIpProtoProtoNumberProgressionTrackingOperOper(res)
		d.Set("oper", DdosDstZoneIpProtoProtoNumberProgressionTrackingOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneIpProtoProtoNumberProgressionTrackingOperOper(ret edpt.DataDdosDstZoneIpProtoProtoNumberProgressionTrackingOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators":           setSliceDdosDstZoneIpProtoProtoNumberProgressionTrackingOperOperIndicators(ret.DtDdosDstZoneIpProtoProtoNumberProgressionTrackingOper.Oper.Indicators),
			"learning_details":     ret.DtDdosDstZoneIpProtoProtoNumberProgressionTrackingOper.Oper.LearningDetails,
			"learning_brief":       ret.DtDdosDstZoneIpProtoProtoNumberProgressionTrackingOper.Oper.LearningBrief,
			"recommended_template": ret.DtDdosDstZoneIpProtoProtoNumberProgressionTrackingOper.Oper.RecommendedTemplate,
			"template_debug_table": ret.DtDdosDstZoneIpProtoProtoNumberProgressionTrackingOper.Oper.TemplateDebugTable,
		},
	}
}

func setSliceDdosDstZoneIpProtoProtoNumberProgressionTrackingOperOperIndicators(d []edpt.DdosDstZoneIpProtoProtoNumberProgressionTrackingOperOperIndicators) []map[string]interface{} {
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

func getObjectDdosDstZoneIpProtoProtoNumberProgressionTrackingOperOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberProgressionTrackingOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberProgressionTrackingOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZoneIpProtoProtoNumberProgressionTrackingOperOperIndicators(in["indicators"].([]interface{}))
		ret.LearningDetails = in["learning_details"].(int)
		ret.LearningBrief = in["learning_brief"].(int)
		ret.RecommendedTemplate = in["recommended_template"].(int)
		ret.TemplateDebugTable = in["template_debug_table"].(int)
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberProgressionTrackingOperOperIndicators(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberProgressionTrackingOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberProgressionTrackingOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberProgressionTrackingOperOperIndicators
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

func dataToEndpointDdosDstZoneIpProtoProtoNumberProgressionTrackingOper(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNumberProgressionTrackingOper {
	var ret edpt.DdosDstZoneIpProtoProtoNumberProgressionTrackingOper

	ret.Oper = getObjectDdosDstZoneIpProtoProtoNumberProgressionTrackingOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.ProtocolNum = d.Get("protocol_num").(string)
	return ret
}
