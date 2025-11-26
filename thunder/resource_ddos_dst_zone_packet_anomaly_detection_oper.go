package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePacketAnomalyDetectionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_packet_anomaly_detection_oper`: Operational Status for the object packet-anomaly-detection\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePacketAnomalyDetectionOperRead,

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
									"maximum": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"minimum": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"threshold": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_anomaly": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"data_source": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}

func resourceDdosDstZonePacketAnomalyDetectionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePacketAnomalyDetectionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePacketAnomalyDetectionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePacketAnomalyDetectionOperOper := setObjectDdosDstZonePacketAnomalyDetectionOperOper(res)
		d.Set("oper", DdosDstZonePacketAnomalyDetectionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePacketAnomalyDetectionOperOper(ret edpt.DataDdosDstZonePacketAnomalyDetectionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators":  setSliceDdosDstZonePacketAnomalyDetectionOperOperIndicators(ret.DtDdosDstZonePacketAnomalyDetectionOper.Oper.Indicators),
			"data_source": ret.DtDdosDstZonePacketAnomalyDetectionOper.Oper.DataSource,
		},
	}
}

func setSliceDdosDstZonePacketAnomalyDetectionOperOperIndicators(d []edpt.DdosDstZonePacketAnomalyDetectionOperOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["rate"] = item.Rate
		in["maximum"] = item.Maximum
		in["minimum"] = item.Minimum
		in["threshold"] = item.Threshold
		in["is_anomaly"] = item.IsAnomaly
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZonePacketAnomalyDetectionOperOper(d []interface{}) edpt.DdosDstZonePacketAnomalyDetectionOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePacketAnomalyDetectionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstZonePacketAnomalyDetectionOperOperIndicators(in["indicators"].([]interface{}))
		ret.DataSource = in["data_source"].(string)
	}
	return ret
}

func getSliceDdosDstZonePacketAnomalyDetectionOperOperIndicators(d []interface{}) []edpt.DdosDstZonePacketAnomalyDetectionOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePacketAnomalyDetectionOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePacketAnomalyDetectionOperOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.Maximum = in["maximum"].(string)
		oi.Minimum = in["minimum"].(string)
		oi.Threshold = in["threshold"].(string)
		oi.IsAnomaly = in["is_anomaly"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePacketAnomalyDetectionOper(d *schema.ResourceData) edpt.DdosDstZonePacketAnomalyDetectionOper {
	var ret edpt.DdosDstZonePacketAnomalyDetectionOper

	ret.Oper = getObjectDdosDstZonePacketAnomalyDetectionOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)
	return ret
}
