package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryPortPatternRecognitionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_entry_port_pattern_recognition_oper`: Operational Status for the object pattern-recognition\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstEntryPortPatternRecognitionOperRead,

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
			"port_num": {
				Type: schema.TypeString, Required: true, Description: "PortNum",
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}

func resourceDdosDstEntryPortPatternRecognitionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortPatternRecognitionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortPatternRecognitionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstEntryPortPatternRecognitionOperOper := setObjectDdosDstEntryPortPatternRecognitionOperOper(res)
		d.Set("oper", DdosDstEntryPortPatternRecognitionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstEntryPortPatternRecognitionOperOper(ret edpt.DataDdosDstEntryPortPatternRecognitionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":              ret.DtDdosDstEntryPortPatternRecognitionOper.Oper.State,
			"timestamp":          ret.DtDdosDstEntryPortPatternRecognitionOper.Oper.Timestamp,
			"peace_pkt_count":    ret.DtDdosDstEntryPortPatternRecognitionOper.Oper.PeacePktCount,
			"war_pkt_count":      ret.DtDdosDstEntryPortPatternRecognitionOper.Oper.WarPktCount,
			"war_pkt_percentage": ret.DtDdosDstEntryPortPatternRecognitionOper.Oper.WarPktPercentage,
			"filter_threshold":   ret.DtDdosDstEntryPortPatternRecognitionOper.Oper.FilterThreshold,
			"filter_count":       ret.DtDdosDstEntryPortPatternRecognitionOper.Oper.FilterCount,
			"filter_list":        setSliceDdosDstEntryPortPatternRecognitionOperOperFilterList(ret.DtDdosDstEntryPortPatternRecognitionOper.Oper.FilterList),
		},
	}
}

func setSliceDdosDstEntryPortPatternRecognitionOperOperFilterList(d []edpt.DdosDstEntryPortPatternRecognitionOperOperFilterList) []map[string]interface{} {
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

func getObjectDdosDstEntryPortPatternRecognitionOperOper(d []interface{}) edpt.DdosDstEntryPortPatternRecognitionOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortPatternRecognitionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Timestamp = in["timestamp"].(string)
		ret.PeacePktCount = in["peace_pkt_count"].(int)
		ret.WarPktCount = in["war_pkt_count"].(int)
		ret.WarPktPercentage = in["war_pkt_percentage"].(int)
		ret.FilterThreshold = in["filter_threshold"].(int)
		ret.FilterCount = in["filter_count"].(int)
		ret.FilterList = getSliceDdosDstEntryPortPatternRecognitionOperOperFilterList(in["filter_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryPortPatternRecognitionOperOperFilterList(d []interface{}) []edpt.DdosDstEntryPortPatternRecognitionOperOperFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortPatternRecognitionOperOperFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortPatternRecognitionOperOperFilterList
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

func dataToEndpointDdosDstEntryPortPatternRecognitionOper(d *schema.ResourceData) edpt.DdosDstEntryPortPatternRecognitionOper {
	var ret edpt.DdosDstEntryPortPatternRecognitionOper

	ret.Oper = getObjectDdosDstEntryPortPatternRecognitionOperOper(d.Get("oper").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.PortNum = d.Get("port_num").(string)

	ret.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
