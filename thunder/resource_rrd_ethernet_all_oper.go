package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRrdEthernetAllOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rrd_ethernet_all_oper`: Operational Status for the object ethernet-all\n\n__PLACEHOLDER__",
		ReadContext: resourceRrdEthernetAllOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start_time": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"end_time": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ethernet_statistics": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ethernet_index": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"stat": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"time": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"tx_bits": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceRrdEthernetAllOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRrdEthernetAllOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRrdEthernetAllOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RrdEthernetAllOperOper := setObjectRrdEthernetAllOperOper(res)
		d.Set("oper", RrdEthernetAllOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRrdEthernetAllOperOper(ret edpt.DataRrdEthernetAllOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"start_time":          ret.DtRrdEthernetAllOper.Oper.StartTime,
			"end_time":            ret.DtRrdEthernetAllOper.Oper.EndTime,
			"ethernet_statistics": setSliceRrdEthernetAllOperOperEthernetStatistics(ret.DtRrdEthernetAllOper.Oper.EthernetStatistics),
		},
	}
}

func setSliceRrdEthernetAllOperOperEthernetStatistics(d []edpt.RrdEthernetAllOperOperEthernetStatistics) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ethernet_index"] = item.EthernetIndex
		in["stat"] = setSliceRrdEthernetAllOperOperEthernetStatisticsStat(item.Stat)
		result = append(result, in)
	}
	return result
}

func setSliceRrdEthernetAllOperOperEthernetStatisticsStat(d []edpt.RrdEthernetAllOperOperEthernetStatisticsStat) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["time"] = item.Time
		in["tx_bits"] = item.Tx_bits
		result = append(result, in)
	}
	return result
}

func getObjectRrdEthernetAllOperOper(d []interface{}) edpt.RrdEthernetAllOperOper {

	count1 := len(d)
	var ret edpt.RrdEthernetAllOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StartTime = in["start_time"].(int)
		ret.EndTime = in["end_time"].(int)
		ret.EthernetStatistics = getSliceRrdEthernetAllOperOperEthernetStatistics(in["ethernet_statistics"].([]interface{}))
	}
	return ret
}

func getSliceRrdEthernetAllOperOperEthernetStatistics(d []interface{}) []edpt.RrdEthernetAllOperOperEthernetStatistics {

	count1 := len(d)
	ret := make([]edpt.RrdEthernetAllOperOperEthernetStatistics, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RrdEthernetAllOperOperEthernetStatistics
		oi.EthernetIndex = in["ethernet_index"].(int)
		oi.Stat = getSliceRrdEthernetAllOperOperEthernetStatisticsStat(in["stat"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRrdEthernetAllOperOperEthernetStatisticsStat(d []interface{}) []edpt.RrdEthernetAllOperOperEthernetStatisticsStat {

	count1 := len(d)
	ret := make([]edpt.RrdEthernetAllOperOperEthernetStatisticsStat, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RrdEthernetAllOperOperEthernetStatisticsStat
		oi.Time = in["time"].(int)
		oi.Tx_bits = in["tx_bits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRrdEthernetAllOper(d *schema.ResourceData) edpt.RrdEthernetAllOper {
	var ret edpt.RrdEthernetAllOper

	ret.Oper = getObjectRrdEthernetAllOperOper(d.Get("oper").([]interface{}))
	return ret
}
