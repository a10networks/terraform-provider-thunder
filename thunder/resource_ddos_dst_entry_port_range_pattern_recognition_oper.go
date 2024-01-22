package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryPortRangePatternRecognitionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_entry_port_range_pattern_recognition_oper`: Operational Status for the object pattern-recognition\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstEntryPortRangePatternRecognitionOperRead,

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
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"port_range_end": {
				Type: schema.TypeString, Required: true, Description: "PortRangeEnd",
			},
			"port_range_start": {
				Type: schema.TypeString, Required: true, Description: "PortRangeStart",
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}

func resourceDdosDstEntryPortRangePatternRecognitionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortRangePatternRecognitionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortRangePatternRecognitionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstEntryPortRangePatternRecognitionOperOper := setObjectDdosDstEntryPortRangePatternRecognitionOperOper(res)
		d.Set("oper", DdosDstEntryPortRangePatternRecognitionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstEntryPortRangePatternRecognitionOperOper(ret edpt.DataDdosDstEntryPortRangePatternRecognitionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":              ret.DtDdosDstEntryPortRangePatternRecognitionOper.Oper.State,
			"timestamp":          ret.DtDdosDstEntryPortRangePatternRecognitionOper.Oper.Timestamp,
			"peace_pkt_count":    ret.DtDdosDstEntryPortRangePatternRecognitionOper.Oper.PeacePktCount,
			"war_pkt_count":      ret.DtDdosDstEntryPortRangePatternRecognitionOper.Oper.WarPktCount,
			"war_pkt_percentage": ret.DtDdosDstEntryPortRangePatternRecognitionOper.Oper.WarPktPercentage,
			"filter_threshold":   ret.DtDdosDstEntryPortRangePatternRecognitionOper.Oper.FilterThreshold,
			"filter_count":       ret.DtDdosDstEntryPortRangePatternRecognitionOper.Oper.FilterCount,
			"filter_list":        setSliceDdosDstEntryPortRangePatternRecognitionOperOperFilterList(ret.DtDdosDstEntryPortRangePatternRecognitionOper.Oper.FilterList),
		},
	}
}

func setSliceDdosDstEntryPortRangePatternRecognitionOperOperFilterList(d []edpt.DdosDstEntryPortRangePatternRecognitionOperOperFilterList) []map[string]interface{} {
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

func getObjectDdosDstEntryPortRangePatternRecognitionOperOper(d []interface{}) edpt.DdosDstEntryPortRangePatternRecognitionOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangePatternRecognitionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Timestamp = in["timestamp"].(string)
		ret.PeacePktCount = in["peace_pkt_count"].(int)
		ret.WarPktCount = in["war_pkt_count"].(int)
		ret.WarPktPercentage = in["war_pkt_percentage"].(int)
		ret.FilterThreshold = in["filter_threshold"].(int)
		ret.FilterCount = in["filter_count"].(int)
		ret.FilterList = getSliceDdosDstEntryPortRangePatternRecognitionOperOperFilterList(in["filter_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryPortRangePatternRecognitionOperOperFilterList(d []interface{}) []edpt.DdosDstEntryPortRangePatternRecognitionOperOperFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangePatternRecognitionOperOperFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangePatternRecognitionOperOperFilterList
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

func dataToEndpointDdosDstEntryPortRangePatternRecognitionOper(d *schema.ResourceData) edpt.DdosDstEntryPortRangePatternRecognitionOper {
	var ret edpt.DdosDstEntryPortRangePatternRecognitionOper

	ret.Oper = getObjectDdosDstEntryPortRangePatternRecognitionOperOper(d.Get("oper").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.PortRangeEnd = d.Get("port_range_end").(string)

	ret.PortRangeStart = d.Get("port_range_start").(string)

	ret.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
