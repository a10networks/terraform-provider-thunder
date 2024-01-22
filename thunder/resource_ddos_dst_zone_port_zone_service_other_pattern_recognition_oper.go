package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceOtherPatternRecognitionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_zone_service_other_pattern_recognition_oper`: Operational Status for the object pattern-recognition\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortZoneServiceOtherPatternRecognitionOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"timestamp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"peace_pkt_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"war_pkt_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"war_pkt_percentage": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"filter_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"filter_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"filter_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"processing_unit": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"filter_enabled": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hardware_filter": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"filter_expr": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"filter_desc": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sample_ratio": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
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

func resourceDdosDstZonePortZoneServiceOtherPatternRecognitionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceOtherPatternRecognitionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceOtherPatternRecognitionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortZoneServiceOtherPatternRecognitionOperOper := setObjectDdosDstZonePortZoneServiceOtherPatternRecognitionOperOper(res)
		d.Set("oper", DdosDstZonePortZoneServiceOtherPatternRecognitionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortZoneServiceOtherPatternRecognitionOperOper(ret edpt.DataDdosDstZonePortZoneServiceOtherPatternRecognitionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":              ret.DtDdosDstZonePortZoneServiceOtherPatternRecognitionOper.Oper.State,
			"timestamp":          ret.DtDdosDstZonePortZoneServiceOtherPatternRecognitionOper.Oper.Timestamp,
			"peace_pkt_count":    ret.DtDdosDstZonePortZoneServiceOtherPatternRecognitionOper.Oper.PeacePktCount,
			"war_pkt_count":      ret.DtDdosDstZonePortZoneServiceOtherPatternRecognitionOper.Oper.WarPktCount,
			"war_pkt_percentage": ret.DtDdosDstZonePortZoneServiceOtherPatternRecognitionOper.Oper.WarPktPercentage,
			"filter_threshold":   ret.DtDdosDstZonePortZoneServiceOtherPatternRecognitionOper.Oper.FilterThreshold,
			"filter_count":       ret.DtDdosDstZonePortZoneServiceOtherPatternRecognitionOper.Oper.FilterCount,
			"filter_list":        setSliceDdosDstZonePortZoneServiceOtherPatternRecognitionOperOperFilterList(ret.DtDdosDstZonePortZoneServiceOtherPatternRecognitionOper.Oper.FilterList),
		},
	}
}

func setSliceDdosDstZonePortZoneServiceOtherPatternRecognitionOperOperFilterList(d []edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionOperOperFilterList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["processing_unit"] = item.ProcessingUnit
		in["filter_enabled"] = item.FilterEnabled
		in["hardware_filter"] = item.HardwareFilter
		in["filter_expr"] = item.FilterExpr
		in["filter_desc"] = item.FilterDesc
		in["sample_ratio"] = item.SampleRatio
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZonePortZoneServiceOtherPatternRecognitionOperOper(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Timestamp = in["timestamp"].(string)
		ret.PeacePktCount = in["peace_pkt_count"].(int)
		ret.WarPktCount = in["war_pkt_count"].(int)
		ret.WarPktPercentage = in["war_pkt_percentage"].(int)
		ret.FilterThreshold = in["filter_threshold"].(int)
		ret.FilterCount = in["filter_count"].(int)
		ret.FilterList = getSliceDdosDstZonePortZoneServiceOtherPatternRecognitionOperOperFilterList(in["filter_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherPatternRecognitionOperOperFilterList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionOperOperFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionOperOperFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionOperOperFilterList
		oi.ProcessingUnit = in["processing_unit"].(string)
		oi.FilterEnabled = in["filter_enabled"].(int)
		oi.HardwareFilter = in["hardware_filter"].(int)
		oi.FilterExpr = in["filter_expr"].(string)
		oi.FilterDesc = in["filter_desc"].(string)
		oi.SampleRatio = in["sample_ratio"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceOtherPatternRecognitionOper(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionOper {
	var ret edpt.DdosDstZonePortZoneServiceOtherPatternRecognitionOper

	ret.Oper = getObjectDdosDstZonePortZoneServiceOtherPatternRecognitionOperOper(d.Get("oper").([]interface{}))

	ret.PortOther = d.Get("port_other").(string)

	ret.ZoneName = d.Get("zone_name").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
