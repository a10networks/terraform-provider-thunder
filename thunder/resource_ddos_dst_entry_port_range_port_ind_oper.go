package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryPortRangePortIndOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_entry_port_range_port_ind_oper`: Operational Status for the object port-ind\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstEntryPortRangePortIndOperRead,

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
									"entry_maximum": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"entry_minimum": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"entry_non_zero_minimum": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"entry_average": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"src_maximum": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"detection_data_source": {
							Type: schema.TypeString, Optional: true, Description: "",
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

func resourceDdosDstEntryPortRangePortIndOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryPortRangePortIndOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryPortRangePortIndOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstEntryPortRangePortIndOperOper := setObjectDdosDstEntryPortRangePortIndOperOper(res)
		d.Set("oper", DdosDstEntryPortRangePortIndOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstEntryPortRangePortIndOperOper(ret edpt.DataDdosDstEntryPortRangePortIndOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"indicators":            setSliceDdosDstEntryPortRangePortIndOperOperIndicators(ret.DtDdosDstEntryPortRangePortIndOper.Oper.Indicators),
			"detection_data_source": ret.DtDdosDstEntryPortRangePortIndOper.Oper.DetectionDataSource,
		},
	}
}

func setSliceDdosDstEntryPortRangePortIndOperOperIndicators(d []edpt.DdosDstEntryPortRangePortIndOperOperIndicators) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["indicator_name"] = item.IndicatorName
		in["indicator_index"] = item.IndicatorIndex
		in["rate"] = item.Rate
		in["entry_maximum"] = item.EntryMaximum
		in["entry_minimum"] = item.EntryMinimum
		in["entry_non_zero_minimum"] = item.EntryNonZeroMinimum
		in["entry_average"] = item.EntryAverage
		in["src_maximum"] = item.SrcMaximum
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstEntryPortRangePortIndOperOper(d []interface{}) edpt.DdosDstEntryPortRangePortIndOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryPortRangePortIndOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Indicators = getSliceDdosDstEntryPortRangePortIndOperOperIndicators(in["indicators"].([]interface{}))
		ret.DetectionDataSource = in["detection_data_source"].(string)
	}
	return ret
}

func getSliceDdosDstEntryPortRangePortIndOperOperIndicators(d []interface{}) []edpt.DdosDstEntryPortRangePortIndOperOperIndicators {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryPortRangePortIndOperOperIndicators, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryPortRangePortIndOperOperIndicators
		oi.IndicatorName = in["indicator_name"].(string)
		oi.IndicatorIndex = in["indicator_index"].(int)
		oi.Rate = in["rate"].(string)
		oi.EntryMaximum = in["entry_maximum"].(string)
		oi.EntryMinimum = in["entry_minimum"].(string)
		oi.EntryNonZeroMinimum = in["entry_non_zero_minimum"].(string)
		oi.EntryAverage = in["entry_average"].(string)
		oi.SrcMaximum = in["src_maximum"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstEntryPortRangePortIndOper(d *schema.ResourceData) edpt.DdosDstEntryPortRangePortIndOper {
	var ret edpt.DdosDstEntryPortRangePortIndOper

	ret.Oper = getObjectDdosDstEntryPortRangePortIndOperOper(d.Get("oper").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.PortRangeEnd = d.Get("port_range_end").(string)

	ret.PortRangeStart = d.Get("port_range_start").(string)

	ret.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
