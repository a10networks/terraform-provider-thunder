package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServicePatternRecognitionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_zone_service_pattern_recognition_oper`: Operational Status for the object pattern-recognition\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortZoneServicePatternRecognitionOperRead,

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

func resourceDdosDstZonePortZoneServicePatternRecognitionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServicePatternRecognitionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServicePatternRecognitionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortZoneServicePatternRecognitionOperOper := setObjectDdosDstZonePortZoneServicePatternRecognitionOperOper(res)
		d.Set("oper", DdosDstZonePortZoneServicePatternRecognitionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortZoneServicePatternRecognitionOperOper(ret edpt.DataDdosDstZonePortZoneServicePatternRecognitionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":              ret.DtDdosDstZonePortZoneServicePatternRecognitionOper.Oper.State,
			"timestamp":          ret.DtDdosDstZonePortZoneServicePatternRecognitionOper.Oper.Timestamp,
			"peace_pkt_count":    ret.DtDdosDstZonePortZoneServicePatternRecognitionOper.Oper.PeacePktCount,
			"war_pkt_count":      ret.DtDdosDstZonePortZoneServicePatternRecognitionOper.Oper.WarPktCount,
			"war_pkt_percentage": ret.DtDdosDstZonePortZoneServicePatternRecognitionOper.Oper.WarPktPercentage,
			"filter_threshold":   ret.DtDdosDstZonePortZoneServicePatternRecognitionOper.Oper.FilterThreshold,
			"filter_count":       ret.DtDdosDstZonePortZoneServicePatternRecognitionOper.Oper.FilterCount,
			"filter_list":        setSliceDdosDstZonePortZoneServicePatternRecognitionOperOperFilterList(ret.DtDdosDstZonePortZoneServicePatternRecognitionOper.Oper.FilterList),
		},
	}
}

func setSliceDdosDstZonePortZoneServicePatternRecognitionOperOperFilterList(d []edpt.DdosDstZonePortZoneServicePatternRecognitionOperOperFilterList) []map[string]interface{} {
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

func getObjectDdosDstZonePortZoneServicePatternRecognitionOperOper(d []interface{}) edpt.DdosDstZonePortZoneServicePatternRecognitionOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServicePatternRecognitionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Timestamp = in["timestamp"].(string)
		ret.PeacePktCount = in["peace_pkt_count"].(int)
		ret.WarPktCount = in["war_pkt_count"].(int)
		ret.WarPktPercentage = in["war_pkt_percentage"].(int)
		ret.FilterThreshold = in["filter_threshold"].(int)
		ret.FilterCount = in["filter_count"].(int)
		ret.FilterList = getSliceDdosDstZonePortZoneServicePatternRecognitionOperOperFilterList(in["filter_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortZoneServicePatternRecognitionOperOperFilterList(d []interface{}) []edpt.DdosDstZonePortZoneServicePatternRecognitionOperOperFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServicePatternRecognitionOperOperFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServicePatternRecognitionOperOperFilterList
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

func dataToEndpointDdosDstZonePortZoneServicePatternRecognitionOper(d *schema.ResourceData) edpt.DdosDstZonePortZoneServicePatternRecognitionOper {
	var ret edpt.DdosDstZonePortZoneServicePatternRecognitionOper

	ret.Oper = getObjectDdosDstZonePortZoneServicePatternRecognitionOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.PortNum = d.Get("port_num").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
