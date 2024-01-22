package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryPortRangeProgressionTrackingOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_entry_port_range_progression_tracking_oper`: Operational Status for the object progression-tracking\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstEntryPortRangeProgressionTrackingOperRead,

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

func resourceDdosDstEntryPortRangeProgressionTrackingOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortRangeProgressionTrackingOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortRangeProgressionTrackingOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstEntryPortRangeProgressionTrackingOperOper := setObjectDdosDstEntryPortRangeProgressionTrackingOperOper(res)
		d.Set("oper", DdosDstEntryPortRangeProgressionTrackingOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstEntryPortRangeProgressionTrackingOperOper(ret edpt.DataDdosDstEntryPortRangeProgressionTrackingOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators": setSliceDdosDstEntryPortRangeProgressionTrackingOperOperIndicators(ret.DtDdosDstEntryPortRangeProgressionTrackingOper.Oper.Indicators),
		},
	}
}

func setSliceDdosDstEntryPortRangeProgressionTrackingOperOperIndicators(d []edpt.DdosDstEntryPortRangeProgressionTrackingOperOperIndicators) []map[string]interface{} {
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

func getObjectDdosDstEntryPortRangeProgressionTrackingOperOper(d []interface{}) edpt.DdosDstEntryPortRangeProgressionTrackingOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangeProgressionTrackingOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryPortRangeProgressionTrackingOperOperIndicators(in["indicators"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryPortRangeProgressionTrackingOperOperIndicators(d []interface{}) []edpt.DdosDstEntryPortRangeProgressionTrackingOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangeProgressionTrackingOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangeProgressionTrackingOperOperIndicators
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

func dataToEndpointDdosDstEntryPortRangeProgressionTrackingOper(d *schema.ResourceData) edpt.DdosDstEntryPortRangeProgressionTrackingOper {
	var ret edpt.DdosDstEntryPortRangeProgressionTrackingOper

	ret.Oper = getObjectDdosDstEntryPortRangeProgressionTrackingOperOper(d.Get("oper").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.PortRangeEnd = d.Get("port_range_end").(string)

	ret.PortRangeStart = d.Get("port_range_start").(string)

	ret.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
